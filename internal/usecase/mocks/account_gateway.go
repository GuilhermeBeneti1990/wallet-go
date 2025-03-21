package mocks

import (
	"github.com/GuilhermeBeneti1990/wallet-go/internal/entities"
	"github.com/stretchr/testify/mock"
)

type AccountGatewayMock struct {
	mock.Mock
}

func (m *AccountGatewayMock) Save(account *entities.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

func (m *AccountGatewayMock) FindById(id string) (*entities.Account, error) {
	args := m.Called(id)
	return args.Get(0).(*entities.Account), args.Error(1)
}

func (m *AccountGatewayMock) UpdateBalance(account *entities.Account) error {
	args := m.Called(account)
	return args.Error(0)
}
