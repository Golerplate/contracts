package session_store_svc_v1

import (
	"context"

	session_store_svc_v1_entities "github.com/golerplate/contracts/clients/sess-store-svc/v1/entities"
)

//go:generate mockgen -source interface.go -destination mocks/mock_sess_store_svc.go -package sess_store_svc_v1_mocks
type SessStoreSvc interface {
	CreateSession(ctx context.Context, userID string) (*session_store_svc_v1_entities.Session, error)
	GetSessionByID(ctx context.Context, id string) (*session_store_svc_v1_entities.Session, error)
}
