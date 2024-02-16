package user_store_v1

import (
	"context"

	user_store_entities_v1 "github.com/Golerplate/contracts/clients/user-store/v1/entities"
)

type UserStoreSvc interface {
	CreateUser(ctx context.Context, req *user_store_entities_v1.UserCreate) (*user_store_entities_v1.User, error)
}
