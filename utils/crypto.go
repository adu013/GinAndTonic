package utils

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
)

// GenerateRandomKey creates a secure random string of a specified byte length.
// For example, passing 32 creates a 256-bit key.
func GenerateRandomKey(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	// Encode bytes into safe, readable text string
	return hex.EncodeToString(bytes), nil
}

// SignValue creates a signed string in the format "value.signature"
func SignValue(value, secret string) string {
	// Create a new HMAC hasher using SHA256 and our secret key
	h := hmac.New(sha256.New, []byte(secret))

	// Feed the value (username) into the hasher
	h.Write([]byte(value))

	// Generate the unique digital fingerprint (signature) in hex format
	signature := hex.EncodeToString(h.Sum(nil))

	// Return the original value combined with the signature, split by a dot
	return value + "||" + signature
	// return signature
}

// VerifySignedValue breaks the cookie apart and validates it.
func VerifySignedValue(signedValue, secret string) (string, error) {
	// Split the cookie value at the "." character
	parts := strings.Split(signedValue, "||")
	if len(parts) != 2 {

		fmt.Printf("[GIN & TONIC - debug] Parts are more than two: %d\n", int(len(parts)))

		return "", errors.New("Invalid session format")
	}

	value := parts[0]             // This is the username
	expectedSignatute := parts[1] // This is the signature string sent by the browser

	fmt.Printf("[GIN & TONIC - debug] value: %s\n", value)
	fmt.Printf("[GIN & TONIC - debug] ecpectedSifnature: %s\n", expectedSignatute)

	// Recompute what the signature SHOULD be using our secret key
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(value))
	actualSignature := hex.EncodeToString(h.Sum(nil))

	// Compare them securely to prevent timing attacks.
	// We convert strings to bytes first.
	if subtle.ConstantTimeCompare([]byte(expectedSignatute), []byte(actualSignature)) != 1 {
		return "", errors.New("Session signatute mismatch (Tampering detected).")
	}

	// If they match perfectly, return the safe username
	return value, nil
}
