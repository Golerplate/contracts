package user_store_svc_v1

import (
	"context"

	user_store_svc_entities_v1 "github.com/Golerplate/contracts/clients/user-store-svc/v1/entities"
)

type UserStoreSvc interface {
	CreateUser(ctx context.Context, req *user_store_svc_entities_v1.UserCreate) (*user_store_svc_entities_v1.User, error)
	GetUserByID(ctx context.Context, id string) (*user_store_svc_entities_v1.User, error)
	GetUserByEmail(ctx context.Context, email string) (*user_store_svc_entities_v1.User, error)
	GetUserByUsername(ctx context.Context, username string) (*user_store_svc_entities_v1.User, error)
}
