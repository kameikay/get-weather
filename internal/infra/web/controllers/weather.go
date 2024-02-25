package controllers

import (
	"github.com/go-chi/chi/v5"
	"github.com/kameikay/get-weather/internal/infra/web/handlers"
)

type WeatherController struct {
	router         chi.Router
	weatherHandler *handlers.WeatherHandler
}

func NewWeatherController(
	router chi.Router,
	weatherHandler *handlers.WeatherHandler,
) *WeatherController {
	return &WeatherController{
		router:         router,
		weatherHandler: weatherHandler,
	}
}

func (wc *WeatherController) Route() {
	wc.router.Route("/weather", func(r chi.Router) {
		r.Get("/", wc.weatherHandler.GetWeather)
	})
}
