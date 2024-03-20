package config

import (
	"log"

	"github.com/joho/godotenv"
)

// InitEnvVar initalizes environment variables declared in ../.env
func InitEnvVar() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}
