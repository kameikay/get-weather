package entity

import (
	"errors"
	"regexp"
	"strings"
)

type Weather struct {
	Cep   string  `json:"cep"`
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}

func NewWeather(cep string) *Weather {
	return &Weather{Cep: cep}
}

func (w *Weather) SetTemperature(tempC float64) {
	w.TempC = tempC
	w.TempF = tempC*1.8 + 32
	w.TempK = tempC + 273
}

func (w *Weather) FormatCEP() (string, error) {
	cepRegEx := `^\d{5}-\d{3}$`

	if regexp.MustCompile(cepRegEx).MatchString(w.Cep) {
		return w.Cep, nil
	}

	if len(w.Cep) > 9 {
		return "", errors.New("invalid cep")
	}

	if len(w.Cep) == 8 && !strings.Contains(w.Cep, "-") {
		return w.Cep[:5] + "-" + w.Cep[5:], nil
	}

	return "", errors.New("invalid cep")
}
