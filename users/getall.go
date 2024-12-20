package users

import (
	"PracticeServer/db"
	"PracticeServer/models"
	"encoding/json"
	"net/http"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	users := make([]models.User, 0)
	tx := db.DB.Find(&users)
	if tx.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := Response{Status: "fail", Message: "Failed to get all user: " + tx.Error.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	response := Response{Status: "success", Message: "Users retrieved successfully", Data: users}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
