package main

import (
	"github.com/kameikay/get-weather/configs"
	"github.com/kameikay/get-weather/internal/infra/web/controllers"
	"github.com/kameikay/get-weather/internal/infra/web/handlers"
	"github.com/kameikay/get-weather/internal/infra/web/webserver"
	"github.com/kameikay/get-weather/internal/service"
)

func main() {
	err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	server := webserver.NewWebServer(":3333")

	server.MountMiddlewares()

	viaCepService := service.NewViaCepService()
	weatherApiService := service.NewWeatherApiService()

	weatherHandler := handlers.NewHandler(viaCepService, weatherApiService)
	weatherController := controllers.NewController(server.Router, weatherHandler)
	weatherController.Route()

	server.Start()
}
