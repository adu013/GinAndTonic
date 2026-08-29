package middlewares

import (
	"crypto/subtle"
	"gin-and-tonic/config"
	"gin-and-tonic/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

const csrfCookieName = "csrf_token"
const csrfFormFieldName = "_csrf"

// SetupCSRF provides cookie-based CSRF protection without external packages
func SetupCSRF() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Handle validation for write requests (POST, PUT, PATCH, DELETE)
		if c.Request.Method == "POST" ||
			c.Request.Method == "PUT" ||
			c.Request.Method == "PATCH" ||
			c.Request.Method == "DELETE" {
			// Get token from cookie
			cookieToken, err := c.Cookie(csrfCookieName)
			if err != nil || cookieToken == "" {
				c.String(http.StatusBadRequest, "CSRF token missing from session!")
				c.Abort()
				return
			}

			// Get token from form submit
			formToken := c.PostForm(csrfFormFieldName)
			if formToken == "" {
				c.String(http.StatusBadRequest, "CSRF token missing from form!")
				c.Abort()
				return
			}

			// Securely compare the tokens to prevent timing attacks
			if subtle.ConstantTimeCompare([]byte(cookieToken), []byte(formToken)) != 1 {
				c.String(http.StatusBadRequest, "CSRF token mismatch")
				c.Abort()
				return
			}

			// if Token is valid, proceed to next
			c.Next()
			return
		}

		// Handle token generation for safe request (GET)
		if c.Request.Method == "GET" {
			// Check if they already have a session cookie
			token, err := c.Cookie(csrfCookieName)
			if err != nil || token == "" {
				// Generate a fresh random token if missing
				token, err := utils.GenerateRandomKey(16)
				if err != nil {
					c.String(http.StatusInternalServerError, "Internal security error")
					c.Abort()
					return
				}

				// Save token to an HTTP-Only, SameSite=Strict cookie
				c.SetCookie(
					csrfCookieName,      // Cookie name
					token,               // Token value
					1800,                // TTL (in seconds, 1800s = 30mins)
					"/",                 // Path
					"",                  // Domain (empty means current domain)
					config.IsSecureHttp, // Secure flag (set config variable IsSecureHttp to true when using HTTPS)
					true,                // HttpOnly flag (blocks JavaScript access)
				)
			}

			// Store the token in Gin's context memory so the controller can find it
			c.Set("csrf_token", token)
		}

		c.Next()
	}
}
