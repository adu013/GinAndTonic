package main

import (
	"gin-and-tonic/config"
	"gin-and-tonic/database"
	"gin-and-tonic/migrations"
	"gin-and-tonic/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initilaize the database connection
	database.ConnectDatabase()

	// Run schema updates separately
	migrations.MigrateDB()

	// Initialize Security keys (Generates SessionSecret)
	config.InitSecurity()

	// Create a new router
	router := gin.Default()

	// Define Templates folder
	router.LoadHTMLGlob("views/templates/*")

	// Define assets folder
	router.StaticFS("/assets", http.Dir("./views/assets"))

	// Set up routes for the router
	routes.SetupRoutes(router)

	// Start server
	router.Run(":8085")
}
