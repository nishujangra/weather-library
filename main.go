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
	days := "1"

	weatherClient := weather.NewWeatherClient(client)

	data, err := weatherClient.GetWeather(city, days)

	if err != nil {
		log.Fatalf("Error getting weather data: %s", err)
	}

	for _, d := range *data {
		fmt.Println("-------------------------------------------------")
		fmt.Printf("City: %s\n", city)
		fmt.Printf("Temperature: %.2f°C or %.2f°F\n", d.Temperature, weather.ConvertCeliusToFahrenheit(d.Temperature))
		fmt.Printf("Humidity: %.2f%\n", d.Humidity)
		fmt.Printf("Description: %s\n", d.Descriptipn)
		fmt.Printf("Date of Forecast: %s\n", d.DateOfForecast)
		fmt.Println("-------------------------------------------------")
	}
}
