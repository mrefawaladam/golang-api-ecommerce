package configs

import (
	"ecommerce-api/internal/domain"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

 func GetEnv(key, defaultValue string) string {
    value := os.Getenv(key)
    if value == "" {
        return defaultValue
    }
    return value
}

func InitDB() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB connection error:", err)
	}
	return db
}

func MigrateDB(db *gorm.DB) {
	err := db.AutoMigrate(
		&domain.User{},
		&domain.Product{},
		&domain.Cart{},
		&domain.Order{},
	)
	if err != nil {
		log.Fatal("Failed to migrate:", err)
	}
}