package routes

import (
	"gin-and-tonic/middlewares"

	"github.com/gin-gonic/gin"
)

// SetupRoutes takes the main router and adds all our pages to it
func SetupRoutes(r *gin.Engine) {
	// Send all /api requests to the API router group
	apiGroup := r.Group("/api")
	SetupAPIRoutes(apiGroup)

	// Send all other requests to the Web router group
	// Apply the CSRF protection only to these web routes
	webGroup := r.Group("/")
	webGroup.Use(middlewares.SetupCSRF())
	SetupWebRoutes(webGroup)
}
