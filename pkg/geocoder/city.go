package geocoder

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type GeoClient struct {
	APIkey  string
	BaseUrl string
}

func NewGeoClient(APIkey string, BaseUrl string) *GeoClient {
	return &GeoClient{
		APIkey:  APIkey,
		BaseUrl: BaseUrl,
	}
}

func (gc *GeoClient) GetLatLong(city string) (float64, float64, error) {
	baseUrl := gc.BaseUrl

	reqUrl, err := url.Parse(baseUrl)

	if err != nil {
		return 0, 0, err
	}

	query := reqUrl.Query()
	query.Add("q", city)
	query.Add("appid", gc.APIkey)

	reqUrl.RawQuery = query.Encode()

	response, err := http.Get(reqUrl.String())

	if err != nil {
		return 0, 0, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return 0, 0, nil
	}

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return 0, 0, err
	}

	var data []struct {
		Lat float64 `json:"lat"`
		Lon float64 `json:"lon"`
	}

	err = json.Unmarshal(body, &data)

	if err != nil {
		return 0, 0, err
	}

	lat := data[0].Lat
	lon := data[0].Lon

	return lat, lon, nil
}
