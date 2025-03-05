package gateway

import "github.com/GuilhermeBeneti1990/wallet-go/internal/entities"

type ClientGateway interface {
	Get(id string) (*entities.Client, error)
	Save(client *entities.Client) error
}
