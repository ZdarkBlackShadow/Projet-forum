package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)


// LoadEnv loads environment variables from a .env file.
//
// If the file cannot be loaded, the function logs a fatal error and stops program execution.
// This is typically called at the start of the application to configure environment variables.
func LoadEnv() {
	errLoad := godotenv.Load(".env")
	if errLoad != nil {
		log.Fatalf("Erreur chargement fichier d'environnement - Impossible de lancer le programme \n\t Erreur : %s", errLoad.Error())
	}
}

// GetEnvWithDefault retrieves the value of an environment variable by key.
//
// If the key does not exist in the environment, it returns the provided default value.
//
// Parameters:
//   - key: the name of the environment variable to retrieve.
//   - defaultValue: the value to return if the environment variable is not set.
//
// Returns:
//   - the environment variable's value if set, otherwise the default value.
func GetEnvWithDefault(key, defaultValue string) string {
	envVar, envErr := os.LookupEnv(key)
	if !envErr {
		return defaultValue
	}
	return envVar
}