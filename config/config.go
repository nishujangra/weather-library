package config

import (
	"os"

	"github.com/joho/godotenv"
)

const (
	WeatherApiBaseUrl  = "https://api.openweathermap.org/data/2.5/forecast"
	GeocoderApiBaseUrl = "https://api.opencagedata.com/geo/1.0/direct"
)

type Client struct {
	APIKey      string
	BaseUrl     string
	GeoBaseUrl  string
	DatabaseUrl string
}

func LoadConfig() *Client {
	err := godotenv.Load()

	if err != nil {
		panic("Error loading .env file")
	}

	return &Client{
		APIKey:      os.Getenv("API_KEY"),
		BaseUrl:     WeatherApiBaseUrl,
		GeoBaseUrl:  GeocoderApiBaseUrl,
		DatabaseUrl: os.Getenv("DATABASE_URL"),
	}
}
