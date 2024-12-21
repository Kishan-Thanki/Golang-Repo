package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetPort() string {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port
		log.Println("PORT not set in .env, defaulting to 8080")
	}
	return port
}
