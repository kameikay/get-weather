package usecase

import (
	"context"

	"github.com/kameikay/get-weather/internal/service"
)

type GetTemperaturesUseCase struct {
	viaCepService     service.ViaCepServiceInterface
	weatherApiService service.WeatherApiServiceInterface
}

type GetTemperaturesUseCaseInterface interface {
	Execute(ctx context.Context, cep string) (Response, error)
}

type Response struct {
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}

func NewGetCEPDataUseCase(
	viaCepService service.ViaCepServiceInterface,
	weatherApiService service.WeatherApiServiceInterface,
) *GetTemperaturesUseCase {
	return &GetTemperaturesUseCase{
		viaCepService:     viaCepService,
		weatherApiService: weatherApiService,
	}
}

func (u *GetTemperaturesUseCase) Execute(ctx context.Context, cep string) (Response, error) {
	cepData, err := u.viaCepService.GetCEPData(ctx, cep)
	if err != nil {
		return Response{}, err
	}

	weatherData, err := u.weatherApiService.GetWeatherData(ctx, cepData.Localidade)
	if err != nil {
		return Response{}, err
	}

	tempF := weatherData.Current.TempC*1.8 + 32
	tempK := weatherData.Current.TempC + 273

	return Response{
		TempC: weatherData.Current.TempC,
		TempF: tempF,
		TempK: tempK,
	}, nil

}
