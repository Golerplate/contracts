package user_store_svc_grpc_v1

import (
	"context"
	"net/http"
	"time"

	"github.com/bufbuild/connect-go"
	"github.com/golerplate/contracts/generated/services/user/store/svc/v1/svcv1connect"
	grpc_interceptors "github.com/golerplate/pkg/grpc/interceptors"

	user_store_svc_v1 "github.com/golerplate/contracts/clients/user-store-svc/v1"
)

type UserStoreSvcClient struct {
	client svcv1connect.UserStoreSvcClient
}

func NewUserStoreSvcClient(ctx context.Context, userStoreSvcClientURL string) user_store_svc_v1.UserStoreSvc {
	iceptorsChain := grpc_interceptors.ClientDefaultChain()

	httpClient := &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:        0,
			MaxConnsPerHost:     0,
			MaxIdleConnsPerHost: 0,
		},
		Timeout: 5 * time.Second,
	}

	return &UserStoreSvcClient{
		client: svcv1connect.NewUserStoreSvcClient(httpClient, userStoreSvcClientURL, connect.WithInterceptors(iceptorsChain...)),
	}
}
