package user_store_grpc_v1

import (
	"context"
	"net/http"
	"time"

	user_store_v1 "github.com/Golerplate/contracts/clients/user-store/v1"
	"github.com/Golerplate/contracts/generated/services/user/store/v1/storev1connect"
	grpcmid "github.com/Golerplate/pkg/grpc/interceptors"
	"github.com/bufbuild/connect-go"
)

type UserStoreClient struct {
	client storev1connect.UserStoreSvcClient
}

func NewUserStoreClient(ctx context.Context, userStoreClientURL string) user_store_v1.UserStoreSvc {
	iceptorsChain := grpcmid.ClientDefaultChain()

	httpClient := &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:        0,
			MaxConnsPerHost:     0,
			MaxIdleConnsPerHost: 0,
		},
		Timeout: 5 * time.Second,
	}

	return &UserStoreClient{
		client: storev1connect.NewUserStoreSvcClient(httpClient, userStoreClientURL, connect.WithInterceptors(iceptorsChain...)),
	}
}
