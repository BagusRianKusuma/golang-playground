package usecase

import (
	"weather-app/internal/entities"
	"weather-app/internal/repositories"
)

type WeatherUsecase struct {
	repo *repositories.WeatherRepository
}

func NewWeatherUsecase(repo *repositories.WeatherRepository) *WeatherUsecase {
	return &WeatherUsecase{repo: repo}
}

func (u *WeatherUsecase) GetWeather(city string) (*entities.Weather, error) {
	return u.repo.GetWeather(city)
}
