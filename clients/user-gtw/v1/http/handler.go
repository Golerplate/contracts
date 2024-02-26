package user_gtw_http_v1

import (
	"context"
	"net/http"
	"time"

	user_gtw_v1 "github.com/golerplate/contracts/clients/user-gtw/v1"
)

type UserGtwClient struct {
	client  *http.Client
	baseURL string
}

func NewUserGtwClient(ctx context.Context, userGtwBaseURL string) user_gtw_v1.UserGtw {
	httpClient := &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:        0,
			MaxConnsPerHost:     0,
			MaxIdleConnsPerHost: 0,
		},
		Timeout: 5 * time.Second,
	}

	return &UserGtwClient{
		client:  httpClient,
		baseURL: userGtwBaseURL,
	}
}
