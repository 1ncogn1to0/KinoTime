package db

import (
	"PracticeServer/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func NewDb() {
	//todo: get env from .env
	dsn := "host=localhost user=postgres password=23111974 dbname=Movie port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	// AutoMigrate models
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Error migrating User model: ", err)
	}
	DB = db
	fmt.Println("Database connected successfully!")
}
