package auth

import (
	"PracticeServer/db"
	"PracticeServer/email"
	"PracticeServer/logging"
	"PracticeServer/models"
	"PracticeServer/users"
	"encoding/json"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var requestData models.User
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if requestData.Name == "" || requestData.Email == "" || requestData.Password == "" {
		http.Error(w, "All fields are required", http.StatusBadRequest)
		return
	}

	hashedPassword, err := HashPassword(requestData.Password)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}
	requestData.Password = hashedPassword

	db.DB.Create(&requestData)

	err = email.SendRegistrationEmail(requestData)
	if err != nil {
		http.Error(w, "Failed to send email", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(users.Response{Status: "success", Message: "User registered successfully"})
}

func ConfirmRegistration(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	if token == "" {
		http.Error(w, "Token is required", http.StatusBadRequest)
		return
	}

	var user models.User
	tx := db.DB.Where("confirmation_token = ?", token).First(&user)
	if tx.Error != nil {
		http.Error(w, "Invalid or expired token", http.StatusBadRequest)
		logging.Logger.Error("Token not found", zap.Error(tx.Error))
		return
	}

	user.IsConfirmed = true
	user.ConfirmationToken = nil

	tx = db.DB.Save(&user)
	if tx.Error != nil {
		http.Error(w, "Failed to confirm user", http.StatusInternalServerError)
		logging.Logger.Error("Failed to update user", zap.Error(tx.Error))
		return
	}

	response := users.Response{Status: "success", Message: "User confirmed successfully"}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
