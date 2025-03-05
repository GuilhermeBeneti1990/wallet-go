package createtransaction

import (
	"testing"

	"github.com/GuilhermeBeneti1990/wallet-go/internal/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type TransactionGatewayMock struct {
	mock.Mock
}

func (m *TransactionGatewayMock) Create(transaction *entities.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
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

func TestCreateTransactionUseCase_Execute(t *testing.T) {
	client1, _ := entities.NewClient("John Doe", "j@email.com")
	account1 := entities.NewAccount(client1)
	account1.Credit(1000)

	client2, _ := entities.NewClient("Jaine Doe", "j2@email.com")
	account2 := entities.NewAccount(client2)
	account2.Credit(1000)

	mockAccount := &AccountGatewayMock{}
	mockAccount.On("FindById", "1").Return(account1, nil)
	mockAccount.On("FindById", "2").Return(account2, nil)

	mockTransaction := &TransactionGatewayMock{}
	mockTransaction.On("Create", mock.Anything).Return(nil)

	inputDTO := CreateTransactionInputDTO{
		AccountIDFrom: account1.ID,
		AccountIDTo:   account2.ID,
		Amount:        100,
	}

	uc := NewCreateTransactionUseCase(mockTransaction, mockAccount)
	outputDTO, err := uc.Execute(inputDTO)
	assert.Nil(t, err)
	assert.NotNil(t, outputDTO)
	mockAccount.AssertExpectations(t)
	mockTransaction.AssertExpectations(t)
	mockAccount.AssertNumberOfCalls(t, "FindById", 2)
	mockTransaction.AssertNumberOfCalls(t, "Create", 1)
}
