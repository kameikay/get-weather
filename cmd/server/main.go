package main

import (
	"github.com/kameikay/get-weather/configs"
	"github.com/kameikay/get-weather/internal/infra/web/controllers"
	"github.com/kameikay/get-weather/internal/infra/web/handlers"
	"github.com/kameikay/get-weather/internal/infra/web/webserver"
)

func main() {
	err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	server := webserver.NewWebServer(":3333")

	server.MountMiddlewares()

	weatherHandler := handlers.NewWeatherHandler()
	weatherController := controllers.NewWeatherController(server.Router, weatherHandler)
	weatherController.Route()

	server.Start()
}
