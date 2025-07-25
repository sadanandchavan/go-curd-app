package config

import (
	"log"

	"github.com/sadanandchavan/go-crud-app/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "root:Tiger@3030@tcp(127.0.0.1:3306)/go-lang-db?parseTime=true&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to database:", err)
	}

	// ✅ Make sure this line exists
	err = database.AutoMigrate(&models.Property{})
	if err != nil {
		log.Fatal("❌ Failed to migrate database:", err)
	}

	DB = database
}
