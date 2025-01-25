package main

import (
	"fmt"
	"log"
	"weather-library/config"
	"weather-library/pkg/weather"
)

func main() {
	client := config.LoadConfig()

	if client.APIKey == "" {
		log.Fatal("API key in environment variable is not set")
	}

	city := "Delhi"

	weatherClient := weather.NewWeatherClient(client)

	data, err := weatherClient.GetWeather(city)

	if err != nil {
		log.Fatalf("Error getting weather data: %s", err)
	}

	fmt.Printf("Weather in %s\n\n", city)
	fmt.Printf("Temperature: %.2f°C or %.2f°F\n", data.Temperature, weather.ConvertCeliusToFahrenheit(data.Temperature))
	fmt.Printf("Humidity: %.2f%%\n", data.Humidity)
	fmt.Printf("Description: %s\n", data.Descriptipn)
}
