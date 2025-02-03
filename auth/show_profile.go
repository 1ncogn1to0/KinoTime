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
		http.Error(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
		return
	}

	var user models.User
	if err := db.DB.Where("email = ?", claims.Email).First(&user).Error; err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(users.Response{
		Status:  "success",
		Message: "Profile loaded",
		Data: map[string]interface{}{
			"ID":        user.ID,
			"name":      user.Name,
			"email":     user.Email,
			"role":      user.RoleID,
			"CreatedAt": user.CreatedAt,
			"UpdatedAt": user.UpdatedAt,
		},
	})
}
