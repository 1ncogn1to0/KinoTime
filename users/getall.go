package users

import (
	"PracticeServer/db"
	"PracticeServer/logging"
	"PracticeServer/models"
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
	"strconv" // преобразования строк в числа
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Извлекает параметры запроса из URL
	name := r.URL.Query().Get("name")
	email := r.URL.Query().Get("email")
	pageStr := r.URL.Query().Get("page")
	sortField := r.URL.Query().Get("sort")  // age, name, email
	sortOrder := r.URL.Query().Get("order") // asc, desc

	query := `SELECT * FROM users WHERE 1=1`
	args := []interface{}{}
	// filter
	if name != "" {
		query += " AND name = ?" // SELECT * FROM users WHERE 1=1 AND name = ?
		args = append(args, name)
	}
	if email != "" {
		query += " AND email = ?"
		args = append(args, email)
	}

	//sort
	if sortOrder != "asc" && sortOrder != "desc" {
		w.WriteHeader(http.StatusBadRequest) // 400
		response := Response{Status: "fail", Message: "Wrong sort order"}
		json.NewEncoder(w).Encode(response)
		return
	}
	if sortField != "name" && sortField != "email" {
		w.WriteHeader(http.StatusBadRequest) // 400
		response := Response{Status: "fail", Message: "Wrong sort field"}
		json.NewEncoder(w).Encode(response)
		return
	}
	if sortOrder != "" && sortField != "" {
		query += " ORDER BY " + sortField + " " + sortOrder // ORDER BY name ASC
	}

	//pagination
	limit := 2
	offset := 0 // смещение (с какого места начать выдачу данных).
	page := 1
	if pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil && p > 1 {
			page = p
		}
	}
	offset = (page - 1) * limit
	query += " LIMIT ? OFFSET ?"
	args = append(args, limit, offset)

	var users []models.User                      // срез юзеров
	tx := db.DB.Raw(query, args...).Scan(&users) // превра ? в данные
	if tx.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logging.Logger.Error("ailed to get users:", zap.Error(tx.Error))
		response := Response{Status: "fail", Message: "Failed to get users: " + tx.Error.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Проверяем, найдены ли пользователи
	if len(users) == 0 {
		w.WriteHeader(http.StatusNotFound) // 404
		logging.Logger.Warn("No users found matching the criteria")
		response := Response{Status: "fail", Message: "No users found matching the criteria"}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Если данные найдены, отправляем их клиенту
	logging.Logger.Info("Users retrieved successfully", zap.String("func", "GetAllUsers"))
	response := Response{Status: "success", Message: "Users retrieved successfully", Data: users}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}
