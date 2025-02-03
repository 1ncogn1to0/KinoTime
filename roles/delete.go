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

func DeleteRole(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Extract ID from URL
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Delete the role
	if err := db.DB.Delete(&models.Role{}, id).Error; err != nil {
		http.Error(w, fmt.Sprintf("Failed to delete role: %v", err), http.StatusInternalServerError)
		return
	}

	// Respond with success
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users.Response{Status: "success", Message: "Role deleted successfully"})
}
