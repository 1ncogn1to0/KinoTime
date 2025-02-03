package roles

import (
	"PracticeServer/db"
	"PracticeServer/models"
	"PracticeServer/users"
	"encoding/json"
	"net/http"
)

func CreateRole(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var role models.Role
	if err := json.NewDecoder(r.Body).Decode(&role); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	db.DB.Create(&role)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(users.Response{Status: "success", Message: "Role created successfully"})
}
