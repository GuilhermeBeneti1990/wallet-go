package gateway

import "github.com/GuilhermeBeneti1990/wallet-go/internal/entities"

type TransactionGateway interface {
	Create(transaction *entities.Transaction) error
}
