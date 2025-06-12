package utils

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"os"
)


// generateSalt generates a random salt of the specified byte size.
//
// Parameters:
//   - size: number of random bytes to generate.
//
// Returns:
//   - a base64-encoded salt string.
//   - an error if random byte generation fails.
func generateSalt(size int) (string, error) {
	bytes := make([]byte, size)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(bytes), nil
}

// HashPassword hashes a password using SHA-512, a randomly generated salt, and a pepper.
//
// The pepper must be provided via the "PEPPER" environment variable.
// The format of the final hash is: SHA-512(password + salt + pepper)
//
// Returns:
//   - the base64-encoded hash.
//   - the salt used during hashing.
//   - an error if PEPPER is missing or any step fails.
func HashPassword(password string) (string, string, error) {
	pepper := os.Getenv("PEPPER")
	if pepper == "" {
		return "", "", fmt.Errorf("PEPPER non défini")
	}

	salt, err := generateSalt(16)
	if err != nil {
		return "", "", err
	}

	data := password + salt + pepper
	hash := sha512.Sum512([]byte(data))
	hashBase64 := base64.StdEncoding.EncodeToString(hash[:])

	return hashBase64, salt, nil
}

// HashPasswordWithSalt hashes a password using SHA-512 with an existing salt and a pepper.
//
// Useful for verifying passwords against stored hashes.
// The pepper must be provided via the "PEPPER" environment variable.
//
// Returns:
//   - the base64-encoded hash.
//   - an error if PEPPER is missing.
func HashPasswordWithSalt(password, salt string) (string, error) {
	pepper := os.Getenv("PEPPER")
	if pepper == "" {
		return "", fmt.Errorf("PEPPER non défini")
	}

	data := password + salt + pepper
	hash := sha512.Sum512([]byte(data))
	hashBase64 := base64.StdEncoding.EncodeToString(hash[:])

	return hashBase64, nil
}