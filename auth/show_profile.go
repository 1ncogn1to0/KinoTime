package auth

import (
	"PracticeServer/db"
	"PracticeServer/models"
	"PracticeServer/users"
	"encoding/json"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
	"time"
)

type UserResponse struct {
	ID        int64       `json:"id"`
	Name      string      `json:"name"`
	Email     string      `json:"email"`
	RoleID    uint        `json:"role_id"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	Role      models.Role `json:"role"`
}

func ShowProfile(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if !strings.HasPrefix(authHeader, "Bearer ") {
		http.Error(w, "Unauthorized: Invalid token format", http.StatusUnauthorized)
		return
	}
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if err != nil || !token.Valid {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Fetch user by email
	var user models.User
	if err := db.DB.Where("email = ?", claims.Email).First(&user).Error; err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	// Fetch role from the roles table
	var role models.Role
	if err := db.DB.Where("id = ?", user.RoleID).First(&role).Error; err != nil {
		http.Error(w, "Role not found", http.StatusForbidden)
		return
	}
	json.NewEncoder(w).Encode(users.Response{Status: "success", Message: "This is your account", Data: UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		RoleID:    user.RoleID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Role:      role,
	}})
}
