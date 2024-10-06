package utils

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

// LoadEnv loads the value of an environment variable from a .env file.
// It reads the .env file specified by the ENV_FILE environment variable,
// retrieves the value for the given key, and returns it. 
// If there's an error loading the .env file or if the variable is not set, 
// it logs an error and exits the application.
//
// Parameters:
// - key: The name of the environment variable.
//
// Returns:
// - The value of the environment variable.
//
// Example:
//   dbHost := utils.LoadEnv("DB_HOST")
func LoadEnv(key string) string {
    envFile := os.Getenv("ENV_FILE")

    if err := godotenv.Load(envFile); err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    value := os.Getenv(key)
    if len(value) == 0 {
        log.Fatalf("Environment variable %s not set or empty", key)
    }

    return value
}
