package handlers

import (
	"log"
	"net/http"

	usecase "github.com/kameikay/get-weather/internal/usecase/weather"
	"github.com/kameikay/get-weather/pkg/exceptions"
	"github.com/kameikay/get-weather/pkg/utils"
)

type WeatherHandler struct {
}

func NewWeatherHandler() *WeatherHandler {
	return &WeatherHandler{}
}

func (h *WeatherHandler) GetWeather(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.JsonResponse(w, utils.ResponseDTO{
			StatusCode: http.StatusMethodNotAllowed,
			Message:    http.StatusText(http.StatusMethodNotAllowed),
			Success:    false,
		})
		return
	}

	cepParam := r.URL.Query().Get("cep")

	if cepParam == "" {
		utils.JsonResponse(w, utils.ResponseDTO{
			StatusCode: http.StatusBadRequest,
			Message:    "cep is required",
			Success:    false,
		})
	}

	getCEPDataUseCase := usecase.NewGetCEPDataUseCase()
	cepData, err := getCEPDataUseCase.Execute(cepParam)
	if err != nil {
		log.Printf("error on get cep data: %v", err)
		if err == exceptions.ErrInvalidCEP {
			utils.JsonResponse(w, utils.ResponseDTO{
				StatusCode: http.StatusUnprocessableEntity,
				Message:    err.Error(),
				Success:    false,
			})
			return
		}

		utils.JsonResponse(w, utils.ResponseDTO{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Success:    false,
		})
		return
	}

	if cepData.Cep == "" {
		utils.JsonResponse(w, utils.ResponseDTO{
			StatusCode: http.StatusNotFound,
			Message:    exceptions.ErrCannotFindZipcode.Error(),
			Success:    false,
		})
		return
	}

	getWeatherDataUseCase := usecase.NewGetWeatherDataUseCase()
	weatherData, err := getWeatherDataUseCase.Execute(cepData.Cep, cepData.Localidade)
	if err != nil {
		log.Printf("error on get weather data: %v", err)
		utils.JsonResponse(w, utils.ResponseDTO{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Success:    false,
		})
		return
	}

	utils.JsonResponse(w, utils.ResponseDTO{
		StatusCode: http.StatusOK,
		Message:    http.StatusText(http.StatusOK),
		Success:    true,
		Data:       weatherData,
	})
}
