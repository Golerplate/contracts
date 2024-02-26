package user_gtw_v1

import (
	"context"

	user_gtw_v1_entities "github.com/golerplate/contracts/clients/user-gtw/v1/entities"
)

type UserGtw interface {
	CreateUser(ctx context.Context, req *user_gtw_v1_entities.CreateUserRequest) error
}
