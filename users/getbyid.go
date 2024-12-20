package users

import (
	"PracticeServer/db"
	"PracticeServer/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func GetByIdUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]

	var user models.User
	tx := db.DB.First(&user, id)
	if tx.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := Response{Status: "fail", Message: "Failed to get user: " + tx.Error.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	response := Response{Status: "success", Message: "successfully retrieved a user", Data: user}
	json.NewEncoder(w).Encode(response)
	return
}
