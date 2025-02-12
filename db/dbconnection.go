package db

import (
	"PracticeServer/models"
	"fmt"
	"github.com/bxcodec/faker/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"os"
	"time"
)

var DB *gorm.DB

type DbConfig struct {
	Host     string `env:"host"`
	User     string `env:"user"`
	Password string `env:"password"`
	Dbname   string `env:"dbname"`
	Port     string `env:"port"`
	Sslmode  string `env:"sslmode"`
}

func NewDb(dbConfig DbConfig) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", dbConfig.Host, dbConfig.User, dbConfig.Password, dbConfig.Dbname, dbConfig.Port, dbConfig.Sslmode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	// AutoMigrate models
	if err := db.AutoMigrate(&models.User{}, &models.Movie{}, &models.Role{}); err != nil {
		log.Fatal("Error migrating User model: ", err)
	}

	// Seed data
	//seedUsers(db)
	//seedRoles(db)

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

func seedRoles(db *gorm.DB) {
	roles := []models.Role{
		{Name: "Администратор", Code: "admin"},
		{Name: "Пользователь", Code: "user"},
		{Name: "Контент-мейкер", Code: "content_maker"},
	}

	for _, role := range roles {
		var existingRole models.Role
		result := db.Where("code = ?", role.Code).First(&existingRole)
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				db.Create(&role)
				fmt.Printf("Role '%s' seeded\n", role.Name)
			} else {
				log.Println("Error checking existing role:", result.Error)
			}
		}
	}
}

func randomBirthDate() time.Time {
	// Генерирует случайную дату рождения от 1970 до 2005 года
	year := rand.Intn(2005-1970) + 1970
	month := time.Month(rand.Intn(12) + 1)
	day := rand.Intn(28) + 1 // Упрощение: все месяцы до 28 дней
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

func LoadDbConfigFromEnv() DbConfig {
	return DbConfig{
		Host:     os.Getenv("host"),
		User:     os.Getenv("user"),
		Password: os.Getenv("password"),
		Dbname:   os.Getenv("dbname"),
		Port:     os.Getenv("port"),
		Sslmode:  os.Getenv("sslmode"),
	}
}
