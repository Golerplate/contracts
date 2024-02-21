package ntfr_email_sdr_grpc_v1

import (
	"context"
	"net/http"
	"time"

	"github.com/bufbuild/connect-go"
	ntfr_email_sdr_v1 "github.com/golerplate/contracts/clients/ntfr-email-sdr/v1"
	"github.com/golerplate/contracts/generated/services/ntfr/email/sdr/v1/sdrv1connect"
	grpc_interceptors "github.com/golerplate/pkg/grpc/interceptors"
)

type NtfrEmailSdrClient struct {
	client sdrv1connect.NtfrEmailSdrClient
}

func NewNtfrEmailSdrClient(ctx context.Context, notifierEmailSdrClientURL string) ntfr_email_sdr_v1.NtfrEmailSdr {
	iceptorsChain := grpc_interceptors.ClientDefaultChain()

	httpClient := &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:        0,
			MaxConnsPerHost:     0,
			MaxIdleConnsPerHost: 0,
		},
		Timeout: 5 * time.Second,
	}

	return &NtfrEmailSdrClient{
		client: sdrv1connect.NewNtfrEmailSdrClient(httpClient, notifierEmailSdrClientURL, connect.WithInterceptors(iceptorsChain...)),
	}
}
