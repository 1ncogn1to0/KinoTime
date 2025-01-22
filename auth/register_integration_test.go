package auth_test

import (
	"PracticeServer/auth"
	"PracticeServer/db"
	"PracticeServer/models"
	"bytes"
	"encoding/json"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func init() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func TestRegisterIntegration(t *testing.T) {
	dbConfig := db.LoadDbConfigFromEnv()
	db.NewDb(dbConfig)

	requestData := models.User{
		Name:     "John Doe",
		Email:    "johndoe@example.com",
		Password: "password123",
		RoleID:   1,
	}

	requestBody, _ := json.Marshal(requestData)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(auth.Register)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, rr.Code)
	}

	var user models.User
	result := db.DB.Where("email = ?", "johndoe@example.com").First(&user)
	if result.Error != nil {
		t.Errorf("User was not inserted into the database")
	}
}
