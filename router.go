package main

import (
	"PracticeServer/email"
	"PracticeServer/middleware"
	"PracticeServer/users"
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", get).Methods("GET")
	router.HandleFunc("/", post).Methods("POST")
	router.HandleFunc("/users", users.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", users.GetByIdUsers).Methods("GET")
	router.HandleFunc("/users/{id}", users.UpdateUsers).Methods("PUT")
	router.HandleFunc("/users/{id}", users.DeleteUsers).Methods("DELETE")
	router.HandleFunc("/users", users.GetAllUsers).Methods("GET")

	router.HandleFunc("/email", email.SendEmail).Methods("POST")
	router.HandleFunc("/send-email", email.SendEmail)
	router.HandleFunc("/index", serveHTML("static/index.html"))

	//middleware only here!
	router.Use(middleware.RateLimit)

	return router
}
func serveHTML(filePath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filePath)
	}
}
