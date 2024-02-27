package usecase

import (
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
)

type GetWeatherDataTestSuite struct {
	suite.Suite
	ctrl *gomock.Controller
}

func TestGetWeatherDataStart(t *testing.T) {
	suite.Run(t, new(GetWeatherDataTestSuite))
}

func (suite *GetWeatherDataTestSuite) GetWeatherDataTestSuiteDown() {
	defer suite.ctrl.Finish()
}

func (suite *GetWeatherDataTestSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())

	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath("../../..")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading config file", err)
	}

	WEATHER_API_KEY := viper.GetString("WEATHER_API_KEY")
	viper.Set("WEATHER_API_KEY", WEATHER_API_KEY)
}

func (suite *GetWeatherDataTestSuite) TestNewGetWeatherDataUseCase() {
	uc := NewGetWeatherDataUseCase()
	suite.NotNil(uc)
}

func (suite *GetWeatherDataTestSuite) TestGetWeatherDataUseCase_Execute() {
	testCases := []struct {
		name        string
		cep         string
		city        string
		expectedErr error
	}{
		{
			name:        "should return a valid response",
			cep:         "80010010",
			city:        "Curitiba",
			expectedErr: nil,
		},
		{
			name:        "should return an empty response (not found)",
			cep:         "00000000",
			city:        "",
			expectedErr: errors.New("cannot find weather data"),
		},
	}

	for _, tc := range testCases {
		suite.T().Run(tc.name, func(t *testing.T) {
			uc := NewGetWeatherDataUseCase()
			_, err := uc.Execute(tc.cep, tc.city)

			suite.Equal(tc.expectedErr, err)
		})
	}
}
