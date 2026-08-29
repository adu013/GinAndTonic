package middlewares

import (
	"fmt"
	"gin-and-tonic/config"
	"gin-and-tonic/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RequireAuth blocks non-logged-in or tampered sessions from private pages
func RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Try to grab the "user_session" cookie from the user's browser
		cookieValue, err := c.Cookie(sessionCookieName)

		// If the cookie doesn't exist, block them and send them to the login page (TODO:)
		if err != nil || cookieValue == "" {
			fmt.Print("[GAT] error: some error or no cookie.\n\n")
			c.Redirect(http.StatusSeeOther, "/")
			c.Abort()
			return
		}

		// Use the helper function to verify the cookie's signature
		// Pass the raw cookie value and the session secret key
		username, err := utils.VerifySignedValue(cookieValue, config.SessionSecret)
		if err != nil {
			// If err is NOT nil, it means someone tampered with the text.
			// Instantly delete their corrupted cookie and kick them out
			c.SetCookie(sessionCookieName, "", -1, "", "", false, true)
			c.Redirect(http.StatusSeeOther, config.LoginRedirectPath)
			c.Abort()
			return
		}

		// If everything is valid, save the username & user in Gin's context
		c.Set("username", username)
		// TODO: Also set user

		// Pass on to next hendler/controller
		c.Next()

	}
}
