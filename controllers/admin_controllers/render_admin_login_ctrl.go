package admincontrollers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ShowLogin displays the secure admin entry page
func RenderAdminLogin(c *gin.Context) {
	csrf_token, exists := c.Get("csrf_token")
	if !exists {
		csrf_token = ""
	}
	c.HTML(http.StatusOK, "admin_login.html", gin.H{
		"Title": "Gin & Tonic Admin",
		"csrf":  csrf_token,
	})
}
