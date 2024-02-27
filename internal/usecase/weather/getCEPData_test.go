package usecase

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kameikay/get-weather/internal/dtos"
	"github.com/kameikay/get-weather/pkg/exceptions"
	"github.com/stretchr/testify/suite"
)

type GetCEPDataTestSuite struct {
	suite.Suite
	ctrl *gomock.Controller
}

func TestGetCEPDataStart(t *testing.T) {
	suite.Run(t, new(GetCEPDataTestSuite))
}

func (suite *GetCEPDataTestSuite) GetCEPDataTestSuiteDown() {
	defer suite.ctrl.Finish()
}

func (suite *GetCEPDataTestSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())
}

func (suite *GetCEPDataTestSuite) TestNewGetCEPDataUseCase() {
	uc := NewGetCEPDataUseCase()
	suite.NotNil(uc)
}

func (suite *GetCEPDataTestSuite) TestGetCEPDataUseCase_Execute() {
	testCases := []struct {
		name        string
		cep         string
		expected    dtos.ViaCEPResponse
		expectedErr error
	}{
		{
			name: "should return a valid response",
			cep:  "01001000",
			expected: dtos.ViaCEPResponse{
				Cep:         "01001-000",
				Logradouro:  "Praça da Sé",
				Complemento: "lado ímpar",
				Bairro:      "Sé",
				Localidade:  "São Paulo",
				Uf:          "SP",
				Ibge:        "3550308",
				Gia:         "1004",
				Ddd:         "11",
				Siafi:       "7107",
			},
			expectedErr: nil,
		},
		{
			name:        "should return an empty response (not found)",
			cep:         "00000000",
			expected:    dtos.ViaCEPResponse{},
			expectedErr: nil,
		},
		{
			name:        "should return an error invalid zip code",
			cep:         "123abc",
			expected:    dtos.ViaCEPResponse{},
			expectedErr: exceptions.ErrInvalidCEP,
		},
	}

	for _, tc := range testCases {
		suite.T().Run(tc.name, func(t *testing.T) {
			uc := NewGetCEPDataUseCase()
			res, err := uc.Execute(tc.cep)

			suite.Equal(tc.expected, res)
			suite.Equal(tc.expectedErr, err)
		})
	}
}
