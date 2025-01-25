package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"weather-library/config"
	"weather-library/pkg/geocoder"
)

type WeatherClient struct {
	APIkey  string
	BaseUrl string
}

type Weather struct {
	Temperature float64 `json:"temp"`
	Humidity    float64 `json:"humidity"`
	Descriptipn string  `json:"description"`
}

func NewWeatherClient(APIkey string, BaseUrl string) *WeatherClient {
	return &WeatherClient{
		APIkey:  APIkey,
		BaseUrl: BaseUrl,
	}
}

func (wc *WeatherClient) GetWeather(city string) (*Weather, error) {
	gcBaseUrl := config.GetGeoBaseUrl()

	gc := geocoder.NewGeoClient(wc.APIkey, gcBaseUrl)
	lat, long, err := gc.GetLatLong(city)

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
	query.Add("cnt", "1")
	query.Add("appid", wc.APIkey)

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
		} `json:"list"`
	}

	if err := json.Unmarshal(body, &parsedData); err != nil {
		return nil, err
	}

	return &Weather{
		Temperature: parsedData.List[0].Main.Temperature,
		Humidity:    parsedData.List[0].Main.Humidity,
		Descriptipn: parsedData.List[0].Weather[0].Description,
	}, nil
}
