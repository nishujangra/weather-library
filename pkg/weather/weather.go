package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"weather-library/config"
	"weather-library/pkg/database"
	"weather-library/pkg/geocoder"
)

type Weather []struct {
	Temperature    float64 `json:"temp"`
	Humidity       float64 `json:"humidity"`
	Descriptipn    string  `json:"description"`
	DateOfForecast string  `json:"dt_txt"`
}

type WeatherClient struct {
	config.Client
}

func NewWeatherClient(client *config.Client) *WeatherClient {
	return &WeatherClient{
		Client: *client,
	}
}

func (wc *WeatherClient) GetWeather(city string, days string) (*Weather, error) {
	gcBaseUrl := wc.GeoBaseUrl

	gc := geocoder.NewGeoClient(wc.APIKey, gcBaseUrl)
	lat, long, err := gc.GetLatLong(city)

	db := database.NewDatabase(wc.DatabaseUrl)
	db.CreateTable()
	defer db.Connection.Close()

	if err != nil {
		return nil, err
	}

	baseUrl := wc.BaseUrl

	reqUrl, err := url.Parse(baseUrl)

	if err != nil {
		return nil, err
	}

	query := reqUrl.Query()
	query.Add("lat", fmt.Sprintf("%.2f", lat))
	query.Add("lon", fmt.Sprintf("%.2f", long))
	query.Add("cnt", days)
	query.Add("appid", wc.APIKey)
	query.Add("units", "metric")

	reqUrl.RawQuery = query.Encode()

	response, err := http.Get(reqUrl.String())

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
		List []struct {
			Main struct {
				Temperature float64 `json:"temp"`
				Humidity    float64 `json:"humidity"`
			} `json:"main"`
			Weather []struct {
				Description string `json:"description"`
			} `json:"weather"`
			DateOfForecast string `json:"dt_txt"`
		} `json:"list"`
	}

	if err := json.Unmarshal(body, &parsedData); err != nil {
		return nil, err
	}

	weather := Weather{}

	for _, data := range parsedData.List {
		db.InsertWeatherData(
			city,
			lat,
			long,
			data.Main.Temperature,
			data.Main.Humidity,
			data.Weather[0].Description,
			data.DateOfForecast,
		)

		weather = append(weather, struct {
			Temperature    float64 `json:"temp"`
			Humidity       float64 `json:"humidity"`
			Descriptipn    string  `json:"description"`
			DateOfForecast string  `json:"dt_txt"`
		}{
			Temperature:    data.Main.Temperature,
			Humidity:       data.Main.Humidity,
			Descriptipn:    data.Weather[0].Description,
			DateOfForecast: data.DateOfForecast,
		})
	}

	return &weather, nil
}
