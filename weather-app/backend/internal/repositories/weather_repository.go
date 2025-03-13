package repositories

import (
	"encoding/json"
	"fmt"
	"net/http"
	"weather-app/internal/entities"
)

type WeatherRepository struct {
	apiKey string
}

func NewWeatherRepository(apiKey string) *WeatherRepository {
	return &WeatherRepository{apiKey: apiKey}
}

func (r *WeatherRepository) GetWeather(city string) (*entities.Weather, error) {
	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s", r.apiKey, city)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var weather entities.Weather
	if err := json.NewDecoder(resp.Body).Decode(&weather); err != nil {
		return nil, err
	}
	return &weather, nil
}
