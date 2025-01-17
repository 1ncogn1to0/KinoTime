package main

import (
	"PracticeServer/auth"
	"PracticeServer/email"
	"PracticeServer/middleware"
	"PracticeServer/roles"
	"PracticeServer/users"
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", get).Methods("GET")
	router.HandleFunc("/", post).Methods("POST")

	authRoutes := router.PathPrefix("/").Subrouter()
	authRoutes.Use(middleware.MiddlewareAuth)
	authRoutes.HandleFunc("/users", users.GetAllUsers).Methods("GET")
	authRoutes.HandleFunc("/users/{id}", users.GetByIdUsers).Methods("GET")
	authRoutes.HandleFunc("/users/{id}", users.UpdateUsers).Methods("PUT")
	authRoutes.HandleFunc("/users/{id}", users.DeleteUsers).Methods("DELETE")

	router.HandleFunc("/email", email.SendEmail).Methods("POST")
	router.HandleFunc("/send-email", email.SendEmail)
	router.HandleFunc("/index", serveHTML("static/index.html"))

	router.HandleFunc("/login", auth.Login).Methods("POST")
	router.HandleFunc("/register", auth.Register).Methods("POST")
	router.HandleFunc("/confirm", auth.ConfirmRegistration).Methods("GET")

	adminRoutes := router.PathPrefix("/admin").Subrouter()
	adminRoutes.Use(middleware.MiddlewareAuth)
	adminRoutes.Use(middleware.MiddlewareRole("admin"))
	adminRoutes.HandleFunc("/roles", roles.GetAllRoles).Methods("GET")
	adminRoutes.HandleFunc("/roles/{id}", roles.GetRoleByID).Methods("GET")
	adminRoutes.HandleFunc("/roles", roles.CreateRole).Methods("POST")
	adminRoutes.HandleFunc("/roles/{id}", roles.UpdateRole).Methods("PUT")
	adminRoutes.HandleFunc("/roles/{id}", roles.DeleteRole).Methods("DELETE")

	//middleware only here!
	router.Use(middleware.RateLimit)

	return router
}
func serveHTML(filePath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filePath)
	}
}
