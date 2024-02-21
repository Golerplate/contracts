package sess_store_svc_grpc_v1

import (
	"context"
	"net/http"
	"time"

	"github.com/bufbuild/connect-go"
	"github.com/golerplate/contracts/generated/services/sess/store/svc/v1/svcv1connect"
	grpc_interceptors "github.com/golerplate/pkg/grpc/interceptors"

	sess_store_svc_v1 "github.com/golerplate/contracts/clients/sess-store-svc/v1"
)

type SessStoreSvcClient struct {
	client svcv1connect.SessStoreSvcClient
}

func NewSessStoreSvcClient(ctx context.Context, sessionStoreSvcClientURL string) sess_store_svc_v1.SessStoreSvc {
	iceptorsChain := grpc_interceptors.ClientDefaultChain()

	httpClient := &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:        0,
			MaxConnsPerHost:     0,
			MaxIdleConnsPerHost: 0,
		},
		Timeout: 5 * time.Second,
	}

	return &SessStoreSvcClient{
		client: svcv1connect.NewSessStoreSvcClient(httpClient, sessionStoreSvcClientURL, connect.WithInterceptors(iceptorsChain...)),
	}
}
