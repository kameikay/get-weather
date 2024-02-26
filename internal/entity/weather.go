package entity

import (
	"errors"
	"strings"
)

type Weather struct {
	Cep                string  `json:"cep"`
	TemperatureCelcius float64 `json:"temp_C"`
}

func NewWeather(cep string) *Weather {
	return &Weather{Cep: cep}
}

func (w *Weather) SetTemperatureCelcius(temp float64) {
	w.TemperatureCelcius = temp
}

func (w *Weather) CalculateFahrenheit() float64 {
	return w.TemperatureCelcius*1.8 + 32
}

func (w *Weather) CalculateKelvin() float64 {
	return w.TemperatureCelcius + 273
}

func (w *Weather) FormatCEP() (string, error) {
	if len(w.Cep) > 9 {
		return "", errors.New("invalid cep")
	}

	if len(w.Cep) == 8 && !strings.Contains(w.Cep, "-") {
		return w.Cep[:5] + "-" + w.Cep[5:], nil
	}

	return "", errors.New("invalid cep")
}
