package gateway

import "github.com/GuilhermeBeneti1990/wallet-go/internal/entities"

type AccountGateway interface {
	Save(account *entities.Account) error
	FindById(id string) (*entities.Account, error)
	UpdateBalance(account *entities.Account) error
}
