package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SessionDestructor() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Tell Gin explicitly how to manage the SameSite parameter
		c.SetSameSite(http.SameSiteLaxMode)

		// Clear the cookie by setting MaxAge to -1
		c.SetCookie(
			sessionCookieName, // Name
			"",                // Clear the value (make it empty text)
			-1,                // MaxAge: -1 tells the browser to delete the cookie instantly!
			"/",               // Path (must match the login path exactly)
			"",                // Domain (blank on local dev)
			false,             // Secure (false for local http testing)
			true,              // HttpOnly
		)
	}
}
