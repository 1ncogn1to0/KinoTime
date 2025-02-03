package roles

import (
	"PracticeServer/db"
	"PracticeServer/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetRoleByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Extract ID from URL
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Find the role
	var role models.Role
	if err := db.DB.First(&role, id).Error; err != nil {
		http.Error(w, fmt.Sprintf("Role not found: %v", err), http.StatusNotFound)
		return
	}

	// Respond with the role
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(role)
}
