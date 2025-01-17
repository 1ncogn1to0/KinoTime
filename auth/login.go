package auth

import (
	"PracticeServer/db"
	"PracticeServer/models"
	"PracticeServer/users"
	"encoding/json"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

var JwtKey = []byte("kino_time_secret")

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var requestData models.User
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	var user models.User
	db.DB.Where("email = ?", requestData.Email).First(&user)
	if user.ID == 0 || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requestData.Password)) != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}
	claims := &Claims{
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users.Response{Status: "success", Message: "Login successful", Data: tokenString})
}
