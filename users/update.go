package users

import (
	"PracticeServer/db"
	"PracticeServer/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func UpdateUsers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	w.Header().Set("Content-Type", "application/json")
	var requestData models.User
	// Парсинг JSON-данных из тела запроса
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		response := Response{Status: "fail", Message: "Некорректное JSON-сообщение"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}
	if requestData.Name == "" || requestData.Email == "" {
		response := Response{Status: "fail", Message: "Поля имя или почта не должны быть пустыми!"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	var user models.User
	tx := db.DB.First(&user, id)

	if tx.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := Response{Status: "fail", Message: "Failed to get user: " + tx.Error.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	user.Name = requestData.Name
	user.Email = requestData.Email

	tx = db.DB.Updates(&user)
	if tx.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := Response{Status: "fail", Message: "Failed to update user: " + tx.Error.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	response := Response{Status: "success", Message: "User updated successfully"}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
	return
}
