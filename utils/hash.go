package utils

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"os"
)

func DisplayPepper() {
	fmt.Println(os.Getenv("PEPPER"))
}

func generateSalt(size int) (string, error) {
	bytes := make([]byte, size)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(bytes), nil
}

func hashPassword(password string) (string, string, error) {
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

func checkPassword(password, salt, expectedHash string) (bool, error) {
	pepper := os.Getenv("PEPPER")
	if pepper == "" {
		return false, fmt.Errorf("PEPPER non défini")
	}

	data := password + salt + pepper
	hash := sha512.Sum512([]byte(data))
	hashBase64 := base64.StdEncoding.EncodeToString(hash[:])

	return hashBase64 == expectedHash, nil
}
