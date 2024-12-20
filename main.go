package main

import (
	"PracticeServer/db"
	"log"
	"net/http"
)

func main() {
	router := NewRouter()
	db.NewDb()
	// Указываем серверу использовать папку "static" для HTML, CSS и JS
	http.Handle("/", http.FileServer(http.Dir("./static")))

	// Запускаем сервер на порту 8080
	log.Println("Сервер запущен: http://localhost:8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("Ошибка запуска сервера:", err)
	}
}
