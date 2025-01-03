package users

import (
	"PracticeServer/db"
	"PracticeServer/logging"
	"PracticeServer/models"
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var requestData models.User
	// Парсинг JSON-данных из тела запроса
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {

		logging.Logger.Error("Некорректное JSON-сообщение", zap.Error(err))

		response := Response{Status: "fail", Message: "Некорректное JSON-сообщение"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}
	if requestData.Name == "" || requestData.Email == "" {

		logging.Logger.Warn("Поля имя или почта не должны быть пустыми!")

		response := Response{Status: "fail", Message: "Поля имя или почта не должны быть пустыми!"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}
	db.DB.Create(&requestData)
	response := Response{Status: "success", Message: "User created successfully"}

	logging.Logger.Info("User created successfully", zap.String("name", requestData.Name), zap.String("func", "CreateUser"))

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
	return
}

type Response struct {
	Status  string      `json:"status"`  // Статус ответа (success/fail)
	Message string      `json:"message"` // Сообщение в ответе
	Data    interface{} `json:"data"`
}
