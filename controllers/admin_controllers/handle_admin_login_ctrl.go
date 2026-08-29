package admincontrollers

import (
	"gin-and-tonic/database"
	"gin-and-tonic/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandleAdminLoginController verifies passwords and sets a simple authentication cookie
func HandleAdminLoginController(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	var user models.User
	// Look up user by email
	if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
		c.HTML(http.StatusUnauthorized, "admin_login.html", gin.H{
			"Title": "Gin & Tonic Admin",
			"Error": "Invalid credentials",
		})
		return
	}

	// Verify the password using the bcrypt helper function inside models/user.go
	if err := user.CheckPassword(password); err != nil {
		c.HTML(http.StatusUnauthorized, "admin_login.html", gin.H{
			"Title": "Gin & Tonic Admin",
			"Error": "Invalid credentials",
		})
		return
	}

	// Verify user is actually an admin
	if user.Role != models.Admin {
		c.HTML(http.StatusForbidden, "admin_login.html", gin.H{
			"Title": "Gin & Tonic Admin",
			"Error": "Access Restricted: Admins Only",
		})
		return
	}

	// Set simple browser cookie token (For template simplicity. In production use JWT or secure sessions!)
	c.SetCookie("admin_session", user.Username, 3600, "/", "", false, true)
	c.Redirect(http.StatusSeeOther, "/admin")
}
