package migrations

import (
	"gin-and-tonic/database"
	"gin-and-tonic/models"
	"log"
)

// MigrateDB updates the database schema to match the Go structs
func MigrateDB() {
	if database.DB == nil {
		log.Fatalf("Cannot migrate: Database connection is not established.")
	}

	log.Println("Running database migrations...")

	// =================================================================
	// Add all your models inside AutoMigrate
	// =================================================================

	// User Model automigration
	err := database.DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Database migration failed: %v", err)
	}

	// Other automigration goes here

	log.Println("Database migration completed successfully!")
}
