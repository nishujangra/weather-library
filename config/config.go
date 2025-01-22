package config

import (
	"os"

	"github.com/joho/godotenv"
)

func GetAPIKey() string {
	err := godotenv.Load()

	if err != nil {
		panic("Error loading .env file")
	}

	return os.Getenv("WEATHER_API_KEY")
}
