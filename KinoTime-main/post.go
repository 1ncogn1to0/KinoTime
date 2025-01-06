package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func post(w http.ResponseWriter, r *http.Request) {
	var requestData RequestData

	// Попытка парсинга JSON-данных из тела запроса
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields() // Отключаем возможность неизвестных полей
	if err := decoder.Decode(&requestData); err != nil {
		response := Response{Status: "fail", Message: "Некорректное JSON-сообщение или неверное название ключа"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	// Проверка наличия поля "message" и его значения
	if requestData.Message == "" {
		// Если message пустое, то success
		response := Response{Status: "success", Message: "Сообщение пустое, но принято"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	// Логирование полученного сообщения
	log.Println("Received message:", requestData.Message)

	// Формирование успешного ответа
	response := Response{Status: "success", Message: "Данные успешно приняты"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
