package email

import (
	"PracticeServer/db"
	"PracticeServer/logging"
	"PracticeServer/models"
	"PracticeServer/users"
	"encoding/json"
	"fmt"
	"github.com/go-gomail/gomail"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func SendEmail(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Failed to parse multipart form data", http.StatusBadRequest)
		return
	}

	jsonBody := r.FormValue("json")
	if jsonBody == "" {
		http.Error(w, "Missing 'json' field in form data", http.StatusBadRequest)
		return
	}

	var emailReq models.Email
	err = json.Unmarshal([]byte(jsonBody), &emailReq)
	if err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	if emailReq.To == "" || emailReq.Subject == "" || emailReq.Body == "" {
		http.Error(w, "Missing required fields: 'to', 'subject', or 'body'", http.StatusBadRequest)
		return
	}

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", os.Getenv("SMTP_USER"))
	mailer.SetHeader("Subject", emailReq.Subject)
	mailer.SetBody("text/plain", emailReq.Body)

	file, fileHeader, err := r.FormFile("file")
	if err == nil {
		defer file.Close()

		fileExt := filepath.Ext(fileHeader.Filename)
		tempFile, err := os.CreateTemp("./image", fmt.Sprintf("upload-*%s", fileExt))
		if err != nil {
			http.Error(w, "Failed to save attachment", http.StatusInternalServerError)
			return
		}

		defer tempFile.Close()

		_, err = io.Copy(tempFile, file)
		if err != nil {
			http.Error(w, "Failed to write attachment to temporary file", http.StatusInternalServerError)
			return
		}

		mailer.Attach(tempFile.Name(), gomail.SetHeader(map[string][]string{
			"Content-Disposition": {fmt.Sprintf(`attachment; filename="%s"`, fileHeader.Filename)},
		}))
	}

	dialer := gomail.NewDialer(
		os.Getenv("SMTP_HOST"),
		587,
		os.Getenv("SMTP_USER"),
		os.Getenv("SMTP_PASS"),
	)

	if emailReq.To != "all" {
		mailer.SetHeader("To", emailReq.To)
		if err := dialer.DialAndSend(mailer); err != nil {
			http.Error(w, fmt.Sprintf("Failed to send email to %s: %v", emailReq.To, err), http.StatusInternalServerError)
			return
		}
	} else {
		var allUsers []models.User
		tx := db.DB.Find(&allUsers)
		if tx.Error != nil {
			http.Error(w, "Failed to retrieve users from database", http.StatusInternalServerError)
			logging.Logger.Error("Failed to get users:", zap.Error(tx.Error))
			return
		}

		for _, user := range allUsers {
			mailer.SetHeader("To", user.Email)
			if err := dialer.DialAndSend(mailer); err != nil {
				logging.Logger.Error("Failed to send email to user", zap.String("email", user.Email), zap.Error(err))
				continue
			}
			logging.Logger.Info("Email sent successfully to user", zap.String("email", user.Email))
		}
	}

	logging.Logger.Info("Email sent successfully", zap.String("To", emailReq.To))
	response := users.Response{Status: "success", Message: "Email sent successfully"}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
