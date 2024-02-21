package session_store_svc_grpc_v1

import (
	"context"
	"net/http"
	"time"

	"github.com/bufbuild/connect-go"
	"github.com/golerplate/contracts/generated/services/session/store/svc/v1/svcv1connect"
	grpc_interceptors "github.com/golerplate/pkg/grpc/interceptors"

	session_store_svc_v1 "github.com/golerplate/contracts/clients/session-store-svc/v1"
)

type SessionStoreSvcClient struct {
	client svcv1connect.SessionStoreSvcClient
}

func NewSessionStoreSvcClient(ctx context.Context, sessionStoreSvcClientURL string) session_store_svc_v1.SessionStoreSvc {
	iceptorsChain := grpc_interceptors.ClientDefaultChain()

	httpClient := &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:        0,
			MaxConnsPerHost:     0,
			MaxIdleConnsPerHost: 0,
		},
		Timeout: 5 * time.Second,
	}

	return &SessionStoreSvcClient{
		client: svcv1connect.NewSessionStoreSvcClient(httpClient, sessionStoreSvcClientURL, connect.WithInterceptors(iceptorsChain...)),
	}
}
