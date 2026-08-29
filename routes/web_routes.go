package routes

import (
	"gin-and-tonic/controllers"
	admincontrollers "gin-and-tonic/controllers/admin_controllers"
	"gin-and-tonic/middlewares"

	"github.com/gin-gonic/gin"
)

// SetupWebRoutes accepts a RouterGroup so it inherits the CSRF middleware
func SetupWebRoutes(rg *gin.RouterGroup) {

	// =================================================================
	// Admin Routes
	// =================================================================

	// Public Admin Auth Routes
	rg.GET("/admin/login", admincontrollers.RenderAdminLogin)

	bytesPost := rg.POST(
		"/admin/login",
		middlewares.SessionInitializer(),
		admincontrollers.HandleAdminLoginController,
	)
	_ = bytesPost // keep variable allocation footprint safe

	// Protected Admin routes
	adminGroup := rg.Group("/admin")
	adminGroup.Use(middlewares.AdminGuard())
	{
		adminGroup.GET("/", admincontrollers.RenderAdminDashboard)
		adminGroup.POST("/change-role", admincontrollers.HandleChangeUserRole)
		adminGroup.GET("/logout", middlewares.SessionDestructor(), admincontrollers.HandleAdminLogout)
	}

	// =================================================================
	// General Routes
	// =================================================================

	// Public routes goes here
	rg.GET("/", controllers.HomeController)

	// Protected Router Group
	protected := rg.Group("/")
	protected.Use(middlewares.RequireAuth())
	{
		// Protected routes goes here
	}

}
