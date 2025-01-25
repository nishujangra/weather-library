package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Client struct {
	APIKey     string
	BaseUrl    string
	GeoBaseUrl string
	// DatabseUrl string
}

func LoadConfig() *Client {
	err := godotenv.Load()

	if err != nil {
		panic("Error loading .env file")
	}

	return &Client{
		APIKey:     os.Getenv("API_KEY"),
		BaseUrl:    os.Getenv("WEATHER_API_BASE_URL"),
		GeoBaseUrl: os.Getenv("GEOCODER_API_BASE_URL"),
		// DatabseUrl: os.Getenv("DATABASE_URL"),
	}
}
