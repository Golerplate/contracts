package user_store_svc_v1

import (
	"context"

	user_store_svc_v1_entities "github.com/golerplate/contracts/clients/user-store-svc/v1/entities"
)

//go:generate mockgen -source interface.go -destination mocks/mock_user_store_svc.go -package user_store_svc_v1_mocks
type UserStoreSvc interface {
	CreateUser(ctx context.Context, req *user_store_svc_v1_entities.UserCreate) (*user_store_svc_v1_entities.User, error)
	UpdateUsername(ctx context.Context, id, username string) (*user_store_svc_v1_entities.User, error)
	GetUserByID(ctx context.Context, id string) (*user_store_svc_v1_entities.User, error)
	GetUserByExternalID(ctx context.Context, externalID string) (*user_store_svc_v1_entities.User, error)
	GetUserByEmail(ctx context.Context, email string) (*user_store_svc_v1_entities.User, error)
	GetUserByUsername(ctx context.Context, username string) (*user_store_svc_v1_entities.User, error)
}
