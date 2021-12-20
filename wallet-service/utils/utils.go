package utils

import (
	"os"

	"github.com/joho/godotenv"
)

// GoDotEnvVariable -> get keys from env file
func GoDotEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	return os.Getenv(key)
}
