package admincontrollers

import (
	"gin-and-tonic/database"
	"gin-and-tonic/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// HandleChangeUserRole updates the user's role permission in the database file
func HandleChangeUserRole(c *gin.Context) {
	userIDStr := c.PostForm("user_id")
	newRole := c.PostForm("role")

	userID, _ := strconv.Atoi(userIDStr)

	var user models.User
	if err := database.DB.First(&user, userID).Error; err == nil {
		// Update the specific role layout field column
		user.Role = models.UserRole(newRole)
		database.DB.Save(&user)
	}

	c.Redirect(http.StatusSeeOther, "/admin")
}
