package config

import (
	"log"
	"os"

	"github.com/sadanandchavan/go-crud-app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// Load from environment variable
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("❌ DATABASE_URL is not set")
	}

	// Connect using PostgreSQL driver
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to PostgreSQL database:", err)
	}

	// Auto-migrate the schema
	err = database.AutoMigrate(&models.Property{})
	if err != nil {
		log.Fatal("❌ Failed to migrate database schema:", err)
	}

	log.Println("✅ Connected to PostgreSQL database successfully.")
	DB = database
}
