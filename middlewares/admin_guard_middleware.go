package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AdminGuard checks if a valid admin session cookie exists before continuing
func AdminGuard() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("admin_session")

		// If no cookie exists, bounce them straight back to the login screen
		if err != nil || cookie == "" {
			c.Redirect(http.StatusSeeOther, "/admin/login")
			c.Abort()
			return
		}

		c.Next()
	}
}
