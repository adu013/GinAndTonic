package middlewares

import (
	"gin-and-tonic/config"
	"gin-and-tonic/utils"

	"github.com/gin-gonic/gin"
)

// IdentifyUser checks for a session to dynamically customize public pages
func IdentifyUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookieValue, err := c.Cookie(sessionCookieName)
		if err == nil && cookieValue != "" {
			// Extract and verify the original username from the signature
			username, err := utils.VerifySignedValue(cookieValue, config.SessionSecret)
			if err == nil {
				c.Set("username", username)
				c.Set("is_authenticated", true)
				c.Next()
				return
			}
		}

		c.Set("username", "")
		c.Set("is_authenticated", false)
		c.Next()

	}
}
