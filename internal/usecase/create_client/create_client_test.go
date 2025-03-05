package createclient

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

func TestCreateClientUseCase_Execute(t *testing.T) {
	t.Run("should create a client", func(t *testing.T) {
		clientGateway := &CliengGatewayMock{}
		clientGateway.On("Save", mock.Anything).Return(nil)

		useCase := NewCreateClientUseCase(clientGateway)

		output, err := useCase.Execute(CreateClientInputDTO{
			Name:  "John Doe",
			Email: "j@email.com",
		})

		assert.Nil(t, err)
		assert.NotNil(t, output)
		assert.NotEmpty(t, output.ID)
		assert.Equal(t, "John Doe", output.Name)
		assert.Equal(t, "j@email.com", output.Email)
		clientGateway.AssertExpectations(t)
		clientGateway.AssertNumberOfCalls(t, "Save", 1)
	})
}
