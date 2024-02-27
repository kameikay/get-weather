package usecase

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/goccy/go-json"
	"github.com/kameikay/get-weather/internal/dtos"
	"github.com/kameikay/get-weather/internal/entity"
	"github.com/spf13/viper"
)

type GetWeatherDataUseCase struct{}

func NewGetWeatherDataUseCase() *GetWeatherDataUseCase {
	return &GetWeatherDataUseCase{}
}

func (uc *GetWeatherDataUseCase) Execute(cep, city string) (dtos.Response, error) {
	WEATHER_API_KEY := viper.GetString("WEATHER_API_KEY")
	formattedCity := strings.ReplaceAll(city, " ", "%20")
	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no", WEATHER_API_KEY, formattedCity)

	res, err := http.Get(url)
	if err != nil {
		return dtos.Response{}, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return dtos.Response{}, errors.New("cannot find weather data")
	}

	var weatherAPIResponse dtos.WeatherAPIResponse
	err = json.NewDecoder(res.Body).Decode(&weatherAPIResponse)
	if err != nil {
		return dtos.Response{}, err
	}

	weather := entity.NewWeather(cep)
	weather.SetTemperature(weatherAPIResponse.Current.TempC)

	return dtos.Response{
		TempC: weather.TempC,
		TempF: weather.TempF,
		TempK: weather.TempK,
	}, nil
}
