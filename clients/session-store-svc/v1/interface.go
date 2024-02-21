package session_store_svc_v1

import (
	"context"

	session_store_svc_v1_entities "github.com/golerplate/contracts/clients/session-store-svc/v1/entities"
)

//go:generate mockgen -source interface.go -destination mocks/mock_session_store_svc.go -package session_store_svc_v1_mocks
type SessionStoreSvc interface {
	CreateSession(ctx context.Context, userID string) (*session_store_svc_v1_entities.Session, error)
	GetSessionByID(ctx context.Context, id string) (*session_store_svc_v1_entities.Session, error)
}
