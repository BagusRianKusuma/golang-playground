package handler

import (
	"encoding/json"
	"net/http"
	"weather-app/internal/usecase"
)

type WeatherHandler struct {
	usecase *usecase.WeatherUsecase
}

func NewWeatherHandler(usecase *usecase.WeatherUsecase) *WeatherHandler {
	return &WeatherHandler{usecase: usecase}
}

func (h *WeatherHandler) GetWeather(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")
	if city == "" {
		http.Error(w, "city parameter is required", http.StatusBadRequest)
		return
	}

	weather, err := h.usecase.GetWeather(city)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(weather)
}
