package users

import (
	"PracticeServer/db"
	"PracticeServer/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func DeleteUsers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if err := db.DB.Delete(&models.User{}, id).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := Response{Status: "fail", Message: "Failed to delete user: " + err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := Response{Status: "success", Message: "User deleted successfully"}
	json.NewEncoder(w).Encode(response)
}
