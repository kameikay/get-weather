package entity

import (
	"testing"

	"github.com/kameikay/get-weather/pkg/exceptions"
	"github.com/stretchr/testify/suite"
)

type WeatherTestSuite struct {
	suite.Suite
}

func TestWeatherStart(t *testing.T) {
	suite.Run(t, new(WeatherTestSuite))
}

func (suite *WeatherTestSuite) TestNewWeather() {
	weather := NewWeather("12345678")
	suite.NotNil(weather)
	suite.Equal(weather.Cep, "12345678")
}

func (suite *WeatherTestSuite) TestWeatherFormatCEP() {
	testCases := []struct {
		name     string
		cep      string
		expected string
		err      error
	}{
		{
			name:     "should return same cep, when cep is valid",
			cep:      "12345-678",
			expected: "12345-678",
			err:      nil,
		},
		{
			name:     "should return formatted cep, when cep is valid",
			cep:      "12345678",
			expected: "12345-678",
			err:      nil,
		},
		{
			name:     "should return error, when cep is invalid",
			cep:      "123456789",
			expected: "",
			err:      exceptions.ErrInvalidCEP,
		},
		{
			name:     "should return error, when cep is invalid",
			cep:      "1234-56781",
			expected: "",
			err:      exceptions.ErrInvalidCEP,
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			weather := NewWeather(tc.cep)
			formatted, err := weather.FormatCEP()
			suite.Equal(formatted, tc.expected)
			suite.Equal(err, tc.err)
		})
	}
}

func (suite *WeatherTestSuite) TestWeatherSetTemperature() {
	weather := NewWeather("12345678")
	weather.SetTemperature(10)
	suite.Equal(weather.TempC, float64(10))
	suite.Equal(weather.TempF, float64(50))
	suite.Equal(weather.TempK, float64(283))
}
