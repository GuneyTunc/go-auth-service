package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv loads environment variables from the specified .env file.
// If the file cannot be loaded, it logs a fatal error and exits.
// This function should typically be called once at the application's startup.
func LoadEnv(filename string) {
	err := godotenv.Load(filename)
	if err != nil {
		log.Fatalf("Error loading %s file: %v", filename, err)
	}
	log.Printf("Successfully loaded environment variables from %s", filename)
}

// GetEnv retrieves the value of an environment variable.
// If the variable is not set, it logs a fatal error and exits.
func GetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("%s environment variable is not set.", key)
	}
	return value
}

// GetEnvOrDefault retrieves the value of an environment variable,
// returning a default value if the variable is not set.
// This is useful for optional configuration values.
func GetEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
