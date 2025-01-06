package main

import (
	"encoding/json"
	"net/http"
)

func get(w http.ResponseWriter, r *http.Request) {
	// Формирование успешного ответа
	response := Response{Status: "success", Message: "GET request successful"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
