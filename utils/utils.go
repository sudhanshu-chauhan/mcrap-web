package utils

import (
	"log"

	"github.com/joho/godotenv"
)

//LoadEnv function loads environment variables from .env file at root folder
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error while loading.env file")
	}
}
