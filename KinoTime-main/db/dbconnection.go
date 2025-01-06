package db

import (
	"PracticeServer/models"
	"fmt"
<<<<<<< HEAD
	"github.com/bxcodec/faker/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"time"
=======
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
>>>>>>> e9dc9fea6c00debc322d4c330c24000cd01592a1
)

var DB *gorm.DB

func NewDb() {
	//todo: get env from .env
<<<<<<< HEAD
	dsn := "host=localhost user=postgres password=23111974 dbname=Movie port=5433 sslmode=prefer"
=======
	dsn := "host=localhost user=postgres password=23111974 dbname=Movie port=5432 sslmode=disable"
>>>>>>> e9dc9fea6c00debc322d4c330c24000cd01592a1
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	// AutoMigrate models
<<<<<<< HEAD
	if err := db.AutoMigrate(&models.User{}, &models.Movie{}); err != nil {
		log.Fatal("Error migrating User model: ", err)
	}

	// Seed data
	seedUsers(db)

=======
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Error migrating User model: ", err)
	}
>>>>>>> e9dc9fea6c00debc322d4c330c24000cd01592a1
	DB = db
	fmt.Println("Database connected successfully!")
}

func CloseDb() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Println("Error retrieving sql.DB from Gorm:", err)
		return
	}

	if err := sqlDB.Close(); err != nil {
		log.Println("Error closing the database connection:", err)
	} else {
		fmt.Println("Database connection closed successfully.")
	}
}
<<<<<<< HEAD

func seedUsers(db *gorm.DB) {
	for i := 0; i < 30; i++ { // Генерация 20 пользователей
		user := models.User{
			Name:      faker.Name(),
			Email:     fmt.Sprintf("%d@astanait.edu.kz", rand.Intn(999999)),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		if err := db.Create(&user).Error; err != nil {
			log.Println("Error seeding user:", err)
		}
	}
	fmt.Println("User seed data added successfully!")
}

func randomBirthDate() time.Time {
	// Генерирует случайную дату рождения от 1970 до 2005 года
	year := rand.Intn(2005-1970) + 1970
	month := time.Month(rand.Intn(12) + 1)
	day := rand.Intn(28) + 1 // Упрощение: все месяцы до 28 дней
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}
=======
>>>>>>> e9dc9fea6c00debc322d4c330c24000cd01592a1
