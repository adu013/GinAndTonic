package config

import (
	"gin-and-tonic/utils"
	"log"
)

// SessionSecret holds our top-secret server key globally
var SessionSecret string

// InitSecurity generates the secret keys needed for the application
func InitSecurity() {
	var err error
	// Generate a cryptographically secure 32-byte random key
	SessionSecret, err = utils.GenerateRandomKey(32)
	if err != nil {
		log.Fatalf("Failed to generate global session secret: %v", err)
	}
}
