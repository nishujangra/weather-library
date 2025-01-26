# Weather Library 🌦️

The Weather Library is a Go-based tool for fetching, displaying, and storing weather data using the OpenWeatherMap API. It leverages PostgreSQL to store weather forecasts for up to the last 16 days, ensuring easy retrieval and efficient management of historical weather data based on city names. The library supports temperature conversion and offers a simple yet extensible architecture

---

## Features ✨
1. Fetch current weather data using city names.
2. Store the last 16 days of forecast data in a PostgreSQL database.
3. Includes temperature (in Celsius and Fahrenheit), humidity, and weather descriptions.
4. Geocoding to fetch latitude and longitude of cities.
5. Extensible and modular design for easy integration.
6. Error handling for API calls and invalid data.

---

## Folder Structure 🗂️
```
.
├── config
│   └── config.go       # Handles environment variables and configuration.
├── pkg
│   └── weather  
│        └── weather.go      # Fetches weather data from the API.
│        └── utils.go        # Includes helper functions (e.g., temperature conversion).
│   └── geocoder
         └── city.go        # Fetches latitude and longitude for a city.
    └── database
         └── database.go        # Store the data of the weather forecast
├── main.go             # Entry point for running the application.
├── .env                # Environment variables file (ignored in version control).
└── README.md           # Documentation for the project.
```

---

## Getting Started 🚀

### Prerequisites

Before running the application, make sure you have the following installed:

- **Go (1.18 or later)**: [Install Go](https://golang.org/doc/install)
- **PostgreSQL**: [Install PostgreSQL](https://www.postgresql.org/download/)
- **Git**: For cloning the repository.
- **.env file**: To store sensitive configuration like API keys and database URLs.

---

### Installation

Follow these steps to run the library on your local machine:

1. **Clone the repository:**
   ```bash
   git clone https://github.com/your-username/weather-library.git
   cd weather-library
   ```

2. **Set up environment variables:**
   Create a `.env` file in the project root and add the following variables:
   ```env
   API_KEY=<your_openweathermap_api_key>
   DATABASE_URL=postgres://<user>:<password>@<host>:<port>/<database>
   ```

   - Replace `your_openweather_api_key` with your actual OpenWeather API key.
   - Replace `username` and `password` in the `DATABASE_URL` with your PostgreSQL credentials.

3. **Install Dependencies**
   Run the following command to install all dependencies:
   ```bash
   go mod tidy
   ```

4. **Set Up PostgreSQL**
   - Start your PostgreSQL server.
   - Create a database named `weatherdb` (or any name of your choice):
     ```sql
     CREATE DATABASE weatherdb;
     ```
   - Ensure the `DATABASE_URL` in the `.env` file matches your database setup.

5. **Run the application:**
   Start the application with:
   ```bash
   go run main.go
   ```

6. **Test the Weather Library**
   Once the application is running, it will:
   - Fetch weather data for a specified city.
   - Display the weather information in the terminal.
   - Store the forecast data in the PostgreSQL database.

---


### How to Use This Library in Your Own Code

To use this library in your Go project, follow these steps:

1. **Add the Weather Library as a Dependency**
   Clone this repository or import it as a module in your project:
   ```bash
   go mod edit -replace weather-library=./path/to/weather-library
   ```

2. **Initialize the Weather Library**
   Import the required packages and set up the configuration:
   ```go
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

   ```

3. **Run Your Project**
   - Make sure you have the `.env` file properly set up and the database configured.
   - Run your project:
     ```bash
     go run main.go
     ```

4. **Expected Output**
   ```bash
   Weather in New York:
   Temperature: 20.00°C (68.00°F)
   Humidity: 50.00%
   Description: clear sky
   Weather data stored in the database.
   ```

---


## How It Works 🛠️

1. **Configuration:**
   The application loads API credentials, base URLs, and the database connection string from the `.env` file using `config/config.go`.

2. **Geocoding:**
   The `pkg/geocoder` package fetches latitude and longitude for the given city using the geocoding API.

3. **Weather Data:**
   - The `pkg/weather/weather.go` file retrieves weather information using the latitude and longitude obtained from the geocoding step.
   - Forecasts are stored in the PostgreSQL database for easy retrieval.

4. **Database Integration:**
   - The `pkg/database/database.go` file manages PostgreSQL connections and CRUD operations.
   - Weather forecasts for the last 16 days are stored in the database, ensuring efficient storage and retrieval.

5. **Utilities:**
   The `pkg/weather/utils.go` file contains helper functions, such as converting temperatures from Celsius to Fahrenheit.

---

## Example Output
Running the application for the city "Delhi" produces the following output:
```plaintext

City: Delhi
Temperature: 27.47°C or 81.45°F
Humidity: 77.00%
Description: scattered clouds
Date of Forecast: 2025-01-26 18:00:00
```

---

## Database Schema 🗄️

The PostgreSQL database uses the following schema to store weather data:

```sql
CREATE TABLE weather_forecast (
    id SERIAL PRIMARY KEY,
    city_name VARCHAR(255) NOT NULL,
    lat FLOAT NOT NULL,
    lon FLOAT NOT NULL,
    temperature FLOAT NOT NULL,
    humidity FLOAT NOT NULL,
    description VARCHAR(255) NOT NULL,
    recorded_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### Key Points:
- `city_name`: Name of the city.
- `lat` and `lon`: Latitude and longitude of the city.
- `temperature`: Temperature in Celsius.
- `humidity`: Humidity percentage.
- `description`: Weather description (e.g., "clear sky").
- `recorded_at`: Timestamp when the data was stored.

---

## Extending the Library 🌟

### Adding a New Feature
To add new functionality, follow the modular structure:
- Create a new package under `/pkg` for your feature.
- Use `config.Client` to access environment variables and API keys.
- Add database operations as needed using `pkg/database/database.go`.

---
### Troubleshooting

1. **Database Connection Error**
   - Ensure PostgreSQL is running and the `DATABASE_URL` in `.env` is correctly set.

2. **Failed to Fetch Weather Data**
   - Verify your API key and ensure the OpenWeather API is reachable.

3. **Table Not Found**
   - If the table is not automatically created, ensure you have database permissions.

4. **Missing Dependencies**
   - Run `go mod tidy` to install any missing dependencies.

---

## Acknowledgments 🙌
- [OpenWeatherMap API](https://openweathermap.org/) for weather data.
- [joho/godotenv](https://github.com/joho/godotenv) for environment variable support.
- PostgreSQL for robust database management.

---

## Contact 📧
For any questions or suggestions, feel free to reach out:
- **Author:** Nishant
- **Email:** ndjangra1027@gmail.com
- **GitHub:** [nishujangra](https://github.com/nishujangra)