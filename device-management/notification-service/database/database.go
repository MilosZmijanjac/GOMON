package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"notification-service/models"
)

func InitDB() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dsn := os.Getenv("DATABASE_DSN")
	if dsn == "" {
		log.Fatalf("DATABASE_DSN environment variable not set")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	return db
}

func MigrateSchemas(db *gorm.DB) {
	if err := db.AutoMigrate(&models.Notification{}); err != nil {
		log.Fatalf("Error migrating schema: %v", err)
	}
}
