package createaccount

import (
	"testing"

	"github.com/GuilhermeBeneti1990/wallet-go/internal/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type CliengGatewayMock struct {
	mock.Mock
}

func (m *CliengGatewayMock) Save(client *entities.Client) error {
	args := m.Called(client)
	return args.Error(0)
}

func (m *CliengGatewayMock) Get(id string) (*entities.Client, error) {
	args := m.Called(id)
	return args.Get(0).(*entities.Client), args.Error(1)
}

type AccountGatewayMock struct {
	mock.Mock
}

func (a *AccountGatewayMock) Save(account *entities.Account) error {
	args := a.Called(account)
	return args.Error(0)
}

func (a *AccountGatewayMock) FindById(id string) (*entities.Account, error) {
	args := a.Called(id)
	return args.Get(0).(*entities.Account), args.Error(1)
}

func TestCreateAccountUseCase_Execute(t *testing.T) {
	t.Run("should create an account", func(t *testing.T) {
		clientGateway := &CliengGatewayMock{}
		clientGateway.On("Get", mock.Anything).Return(&entities.Client{}, nil)

		accountGateway := &AccountGatewayMock{}
		accountGateway.On("Save", mock.Anything).Return(nil)

		useCase := NewCreateAccountUseCase(accountGateway, clientGateway)

		output, err := useCase.Execute(CreateAccountInputDTO{
			ClientID: "1",
		})

		assert.Nil(t, err)
		assert.NotNil(t, output)
		assert.NotEmpty(t, output.ID)
		accountGateway.AssertExpectations(t)
		accountGateway.AssertNumberOfCalls(t, "Save", 1)
	})
}
