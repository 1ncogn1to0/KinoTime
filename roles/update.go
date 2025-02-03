package roles

import (
	"PracticeServer/db"
	"PracticeServer/models"
	"PracticeServer/users"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func UpdateRole(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Extract ID from URL
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Decode the new role details
	var updatedRole models.Role
	if err := json.NewDecoder(r.Body).Decode(&updatedRole); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Update the role
	tx := db.DB.Model(&models.Role{}).Where("id = ?", id).Updates(updatedRole)
	if tx.Error != nil {
		http.Error(w, fmt.Sprintf("Failed to update role: %v", tx.Error), http.StatusInternalServerError)
		return
	}

	// Respond with success
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users.Response{Status: "success", Message: "Role updated successfully"})
}
