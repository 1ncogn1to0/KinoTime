package roles

import (
	"PracticeServer/db"
	"PracticeServer/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetAllRoles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Fetch all roles
	var roles []models.Role
	if err := db.DB.Find(&roles).Error; err != nil {
		http.Error(w, fmt.Sprintf("Failed to fetch roles: %v", err), http.StatusInternalServerError)
		return
	}

	// Respond with the roles
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(roles)
}
