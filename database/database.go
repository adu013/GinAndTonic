package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Global db variable - DB
var DB *gorm.DB

// ConnectDatabase sets the database connection
func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("local.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	DB = database
	log.Println("Database connection successfully established!")
}
