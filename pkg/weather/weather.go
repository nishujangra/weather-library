package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type WeatherClient struct {
	APIkey string
}

type Weather struct {
	Temperature float64 `json:"temp"`
	Humidity    float64 `json:"humidity"`
	Descriptipn string  `json:"description"`
}

func NewWeatherClient(APIkey string) *WeatherClient {
	return &WeatherClient{APIkey: APIkey}
}

func (wc *WeatherClient) GetWeather(city string) (*Weather, error) {
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, wc.APIkey)

	response, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	var parsedData struct {
		Main struct {
			Temperature float64 `json:"temp"`
			Humidity    float64 `json:"humidity"`
		}
		Weather []struct {
			Description string `json:"description"`
		} `json:"weather"`
	}

	if err := json.Unmarshal(body, &parsedData); err != nil {
		return nil, err
	}

	return &Weather{
		Temperature: parsedData.Main.Temperature,
		Humidity:    parsedData.Main.Humidity,
		Descriptipn: parsedData.Weather[0].Description,
	}, nil
}
