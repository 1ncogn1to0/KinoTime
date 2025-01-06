package main

import (
	"PracticeServer/db"
	"PracticeServer/logging"
<<<<<<< HEAD
	"PracticeServer/movies"
=======
>>>>>>> e9dc9fea6c00debc322d4c330c24000cd01592a1
	"context"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	router := NewRouter()
	db.NewDb()
<<<<<<< HEAD
	movies.SeedMovies()
=======
>>>>>>> e9dc9fea6c00debc322d4c330c24000cd01592a1
	// Указываем серверу использовать папку "static" для HTML, CSS и JS
	http.Handle("/", http.FileServer(http.Dir("./static")))

	err := logging.NewLogger()
	if err != nil {
		log.Fatal(err)
	}

	// Запускаем сервер на порту 8080
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	go func() {
		logging.Logger.Info("Сервер запущен: http://localhost:8080/index")
		if err = server.ListenAndServe(); err != nil {
			logging.Logger.Error("Ошибка запуска сервера:", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit
	logging.Logger.Info("Получен сигнал завершения, остановка сервера...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logging.Logger.Error("Ошибка при завершении сервера:", zap.Error(err))
	} else {
		logging.Logger.Info("Сервер успешно остановлен.")
	}

	db.CloseDb()
	logging.Logger.Sync()
}