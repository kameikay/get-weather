package usecase

import (
	"net/http"

	"github.com/goccy/go-json"
	"github.com/kameikay/get-weather/internal/dtos"
	"github.com/kameikay/get-weather/internal/entity"
)

type GetCEPDataUseCase struct {
}

func NewGetCEPDataUseCase() *GetCEPDataUseCase {
	return &GetCEPDataUseCase{}
}

func (u *GetCEPDataUseCase) Execute(cep string) (dtos.ViaCEPResponse, error) {
	weather := entity.NewWeather(cep)
	cepFormatted, err := weather.FormatCEP()
	if err != nil {
		return dtos.ViaCEPResponse{}, err
	}

	var viaCepResponseDTO dtos.ViaCEPResponse
	res, err := http.Get("https://viacep.com.br/ws/" + cepFormatted + "/json/")
	if err != nil {
		return dtos.ViaCEPResponse{}, err
	}

	defer res.Body.Close()

	if res.StatusCode == 200 {
		err = json.NewDecoder(res.Body).Decode(&viaCepResponseDTO)
		if err != nil {
			return dtos.ViaCEPResponse{}, err
		}

		return viaCepResponseDTO, nil
	}

	return dtos.ViaCEPResponse{}, nil
}
