# Weather Library 🌦️

The Weather Library is a Go-based tool for fetching and displaying weather data using the OpenWeatherMap API. It leverages modular design, including geocoding functionality, to provide accurate weather information based on city names. The library supports temperature conversion and offers a simple yet extensible architecture.

---

## Features ✨
- Fetch current weather data using city names.
- Includes temperature (in Celsius and Fahrenheit), humidity, and weather descriptions.
- Geocoding to fetch latitude and longitude of cities.
- Extensible and modular design for easy integration.
- Error handling for API calls and invalid data.

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
├── main.go             # Entry point for running the application.
├── .env                # Environment variables file (ignored in version control).
└── README.md           # Documentation for the project.
```

---

## Getting Started 🚀

### Prerequisites
- Go 1.20 or higher installed on your machine.
- An OpenWeatherMap API key. You can get one by signing up at [OpenWeatherMap](https://openweathermap.org/).

---

### Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/nishujangra/weather-library.git
   cd weather-library
   ```

2. **Set up environment variables:**
   Create a `.env` file in the project root and add the following variables:
   ```env
   API_KEY=<your_openweathermap_api_key>
   WEATHER_API_BASE_URL=https://api.openweathermap.org/data/2.5/forecast
   GEOCODER_API_BASE_URL=http://api.openweathermap.org/geo/1.0/direct
   ```

3. **Run the application:**
   ```bash
   go run main.go
   ```

---

## How It Works 🛠️

1. **Configuration:**
   The application loads API credentials and base URLs from the `.env` file using `config/config.go`.

2. **Geocoding:**
   The `pkg/geocoder/city.go` package fetches latitude and longitude for the given city using the geocoding API.

3. **Weather Data:**
   The `pkg/weather/weather.go` file retrieves weather information using the latitude and longitude obtained from the geocoding step.

4. **Utilities:**
   The `pkg/weather/utils.go` file contains helper functions, such as converting temperatures from Celsius to Fahrenheit.

---

## Example Output
Running the application for the city "Delhi" produces the following output:
```plaintext
Weather in Delhi

Temperature: 30.00°C or 86.00°F
Humidity: 70.00%
Description: clear sky
```

---

## Extending the Library 🌟

### Adding a New Feature
To add new functionality, follow the modular structure:
- Create a new package under `/pkg` for your feature.
- Use `config.Client` to access environment variables and API keys.

### Using in Other Applications
You can integrate the `weather` and `geocoder` packages into other Go projects. Simply import the necessary modules and configure the API keys.

---

## Error Handling ⚠️
The library includes error handling for:
- Missing API keys.
- API connection failures.
- Unexpected responses or status codes.

If an error occurs, the application logs an appropriate message and exits.

---

## Contributing 🤝
Contributions are welcome! Please follow these steps:
1. Fork the repository.
2. Create a new branch: `git checkout -b feature-name`.
3. Commit your changes: `git commit -m 'Add a new feature'`.
4. Push to the branch: `git push origin feature-name`.
5. Submit a pull request.

---

## Acknowledgments 🙌
- [OpenWeatherMap API](https://openweathermap.org/) for weather data.
- [joho/godotenv](https://github.com/joho/godotenv) for environment variable support.

---

## Contact 📧
For any questions or suggestions, feel free to reach out:
- **Author:** Nishant
- **Email:** ndjangra1027@gmail.com
- **GitHub:** [nishujangra](https://github.com/nishujangra)