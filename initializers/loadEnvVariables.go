package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

// Load the environment variables
func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
