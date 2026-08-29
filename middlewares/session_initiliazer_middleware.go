package middlewares

import (
	"gin-and-tonic/config"
	"gin-and-tonic/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SessionInitializer() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the username (Make sure you set this in a previous middleware or login step)
		username := c.PostForm("username") // or c.GetString("username") depending on your setup

		// Verify it isn't empty
		if username != "" {
			log.Println("🏆 Middleware: Username is valid, setting cookie now.")

			// Securely sign the username string using our helper
			token := utils.SignValue(username, config.SessionSecret)

			// Set the modern SameSite standard rule
			c.SetSameSite(http.SameSiteLaxMode)

			// 4. Set the cookie perfectly inside the middleware pipeline
			c.SetCookie(
				sessionCookieName, // Name
				token,             // Value
				3600*24,           // MaxAge (1 day)
				"/",               // Path
				"",                // Domain (Blank on local dev)
				false,             // Secure (False for local http testing)
				true,              // HttpOnly
			)

			// Save it to Gin context so your controller can see it if needed
			c.Set("username", username)

			// Set Gin context "is_authenticated" to true
			c.Set("is_authenticated", true)
		} else {
			log.Println("Middleware: Cookie skipped because username string is empty!")
		}

		// Call Next() to pass control to the next handler/controller
		c.Next()
	}
}
