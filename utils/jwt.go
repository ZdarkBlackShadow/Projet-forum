// Package utils provides utility functions for the forum application.
package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateJWT creates a new JWT token for the given user ID.
//
// It uses the JWT_SECRET environment variable to sign the token.
// The token includes the following claims:
//   - sub: the user ID
//   - exp: expiration time (1 hour from creation)
//   - iat: issued at time
//   - iss: issuer ("forum")
//
// Returns the signed token string and any error encountered.
// If JWT_SECRET is not set, it returns an error.
func GenerateJWT(userID string) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "", fmt.Errorf("JWT_SECRET non défini")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour).Unix(),
		"iat": time.Now().Unix(),
		"iss": "forum",
	})
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// VerifyJWT validates a JWT token and extracts the user ID.
//
// It uses the JWT_SECRET environment variable to verify the token signature.
// The function performs the following validations:
//   - Checks if the token signature is valid
//   - Verifies the signing method is HMAC
//   - Ensures the token has not expired
//
// Returns the user ID from the token's "sub" claim and any error encountered.
// If JWT_SECRET is not set, the token is invalid, or the token has expired,
// it returns an appropriate error.
func VerifyJWT(tokenString string) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "", fmt.Errorf("JWT_SECRET non défini")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("méthode de signature inattendue : %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil || !token.Valid {
		return "", fmt.Errorf("token invalide : %v", err)
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if exp, ok := claims["exp"].(float64); ok {
			if int64(exp) < time.Now().Unix() {
				return "", fmt.Errorf("token expiré")
			}
		}
		userID, _ := claims["sub"].(string)
		return userID, nil
	}

	return "", fmt.Errorf("invalid claims")
}