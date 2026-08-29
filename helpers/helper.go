package helpers

import (
	"gin-and-tonic/utils"
	"log"
)

// GenerateSessionKey generates session key of 32 bytes
func GenerateSessionSecretKey() string {
	sessionKey, err := utils.GenerateRandomKey(32)
	if err != nil {
		log.Fatalf("Failed to generate session key: %v", err)
	}

	return sessionKey
}

// GenerateCSRFKey generates session key of 16 bytes
func GenerateCSRFKey() string {
	csrfKey, err := utils.GenerateRandomKey(16)
	if err != nil {
		log.Fatalf("Failed to generate session key: %v", err)
	}

	return csrfKey
}
