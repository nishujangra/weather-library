# Weather Library 🌦️

A simple and extensible Go library to fetch weather information from APIs like OpenWeatherMap. This library provides utilities to get real-time weather data, such as temperature, humidity, and weather descriptions for any city.

---

## Features
- 🌡️ Fetch current weather data for any city.
- 📊 Supports temperature in Celsius and Farhanitrate both.
- 🔄 Easily extendable for additional APIs or features.
- ⚙️ Includes a CLI example in the main.go file of the root folder for quick usage.

---

## Folder Structure

```
weather-library/
├── main.go      # CLI example for the weather library

├── config/        # Configuration management (e.g., API keys)
│       └── config.go
├── pkg/
│   ├── weather/         # Core weather functionality
│   │   ├── weather.go   # Weather API logic
│   │   └── util.go      # Utility functions
├── .env                 # Environment variables (e.g., API keys)
├── go.mod               # Go module file
└── README.md            # Documentation
```

---

## Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/nishujangra/weather-library.git
   cd weather-library
   ```

2. **Install dependencies:**
   ```bash
   go mod tidy
   ```

3. **Set up your environment:**
   Create a `.env` file in the root of the project with the following content:
   ```plaintext
   WEATHER_API_KEY=your_openweathermap_api_key
   ```

---

## Usage

### Importing the Library
```go
package main

import (
    "fmt"
    "log"

    "github.com/nishujangra/weather-library/pkg/weather"
    "github.com/nishujangra/weather-library/config"
)

func main() {
    apiKey := config.GetAPIKey()
    city := "DELHI"

    service := weather.NewWeatherService(apiKey)
    data, err := service.GetWeather(city)
    if err != nil {
        log.Fatalf("Error fetching weather data: %v", err)
    }

    fmt.Printf("Weather in %s:\nTemperature: %.2f°C\nHumidity: %d%%\nDescription: %s\n", 
        city, data.Temperature, data.Humidity, data.Description)
}
```

### Running the CLI Example

1. Run the application in root directory:
   ```bash
   go run main.go
   ```

2. Example output:
   ```plaintext
   Weather in London:
   Temperature: 15.5°C
   Humidity: 78%
   Description: clear sky
   ```

---

## Utilities

### Convert Temperature
The library includes utility functions for temperature conversion:
```go
celsius := 0.0
fahrenheit := weather.ConvertToFahrenheit(celsius)
fmt.Printf("%.2f°C is %.2f°F\n", celsius, fahrenheit)
```

---

## Future Enhancements
- 🌐 Add support for multiple weather APIs.
- 🗂️ Implement caching for repeated requests.
- 📈 Add historical weather data support.

---

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue.

---


## Acknowledgments

- **OpenWeatherMap API** for providing weather data.
