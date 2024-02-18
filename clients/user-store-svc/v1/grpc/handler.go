package user_store_svc_grpc_v1

import (
	"context"
	"net/http"
	"time"

	user_store_svc_v1 "github.com/Golerplate/contracts/clients/user-store-svc/v1"
	"github.com/Golerplate/contracts/generated/services/user/store/svc/v1/svcv1connect"
	grpcmid "github.com/Golerplate/pkg/grpc/interceptors"
	"github.com/bufbuild/connect-go"
)

type UserStoreSvcClient struct {
	client svcv1connect.UserStoreSvcClient
}

func NewUserStoreSvcClient(ctx context.Context, userStoreSvcClientURL string) user_store_svc_v1.UserStoreSvc {
	iceptorsChain := grpcmid.ClientDefaultChain()

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
