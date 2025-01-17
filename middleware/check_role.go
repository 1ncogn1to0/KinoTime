package middleware

import (
	"PracticeServer/auth"
	"PracticeServer/db"
	"PracticeServer/models"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
)

func MiddlewareRole(requiredRole string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenString := r.Header.Get("Authorization")
			if tokenString == "" {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			claims := &auth.Claims{}
			token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
				return auth.JwtKey, nil
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

			// Check if the role matches the required role
			if role.Name != requiredRole {
				http.Error(w, "Forbidden: insufficient role permissions", http.StatusForbidden)
				return
			}

			// Pass the request to the next handler
			next.ServeHTTP(w, r)
		})
	}
}
