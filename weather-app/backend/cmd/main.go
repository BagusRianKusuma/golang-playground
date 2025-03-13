package main

import (
	"log"
	"weather-app/internal/config"
	"weather-app/internal/handler"
	"weather-app/internal/repositories"
	"weather-app/internal/usecase"
	"weather-app/pkg/http"
)

func main() {
	cfg := config.LoadConfig()

	weatherRepo := repositories.NewWeatherRepository(cfg.WeatherAPIKey)
	weatherUsecase := usecase.NewWeatherUsecase(weatherRepo)
	weatherHandler := handler.NewWeatherHandler(weatherUsecase)

	server := http.NewServer()
	server.RegisterRoutes(weatherHandler)
	log.Fatal(server.Start(":8080"))
}
