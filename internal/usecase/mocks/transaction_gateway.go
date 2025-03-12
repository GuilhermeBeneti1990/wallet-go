package mocks

import (
	"github.com/GuilhermeBeneti1990/wallet-go/internal/entities"
	"github.com/stretchr/testify/mock"
)

type TransactionGatewayMock struct {
	mock.Mock
}

func (m *TransactionGatewayMock) Create(transaction *entities.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}
