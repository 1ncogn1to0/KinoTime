package email

import (
	"PracticeServer/db"
	"PracticeServer/logging"
	"PracticeServer/models"
	"fmt"
	"github.com/go-gomail/gomail"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"os"
)

func SendRegistrationEmail(newUser models.User) error {

	// Validate required fields
	if newUser.Email == "" || newUser.Password == "" {
		return fmt.Errorf("email and password are required")
	}

	// Ensure the primary key (ID) is set
	if newUser.ID == 0 {
		return fmt.Errorf("user ID is required for updates")
	}

	// Generate a confirmation token
	confirmationToken := uuid.NewString()

	newUser.IsConfirmed = false
	newUser.ConfirmationToken = &confirmationToken

	// Update the user in the database
	tx := db.DB.Model(&models.User{}).Where("id = ?", newUser.ID).Updates(&newUser)
	if tx.Error != nil {
		logging.Logger.Error("Failed to update user", zap.Error(tx.Error))
		return fmt.Errorf("failed to update user: %v", tx.Error)
	}

	// Prepare and send the confirmation email
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", os.Getenv("SMTP_USER"))
	mailer.SetHeader("To", newUser.Email)
	mailer.SetHeader("Subject", "Confirm Your Registration")
	confirmationLink := fmt.Sprintf("%s/confirm?token=%s", os.Getenv("APP_URL"), confirmationToken)
	mailer.SetBody("text/plain", fmt.Sprintf("Please confirm your registration by clicking on the following link: %s", confirmationLink))

	// Configure the email dialer
	dialer := gomail.NewDialer(
		os.Getenv("SMTP_HOST"),
		587, // Adjust the port if needed
		os.Getenv("SMTP_USER"),
		os.Getenv("SMTP_PASS"),
	)

	// Send the email
	if err := dialer.DialAndSend(mailer); err != nil {
		logging.Logger.Error("Failed to send email", zap.Error(err))
		return fmt.Errorf("failed to send email: %v", err)
	}

	return nil
}
