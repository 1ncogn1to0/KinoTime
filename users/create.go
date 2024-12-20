package users

import (
	"PracticeServer/db"
	"PracticeServer/models"
	"encoding/json"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
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
	db.DB.Create(&requestData)
	response := Response{Status: "success", Message: "User created successfully"}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
	return
}

type Response struct {
	Status  string      `json:"status"`  // Статус ответа (success/fail)
	Message string      `json:"message"` // Сообщение в ответе
	Data    interface{} `json:"data"`
}
