package notifier_email_sdr_grpc_v1

import (
	"context"
	"net/http"
	"time"

	"github.com/bufbuild/connect-go"
	notifier_email_sdr_v1 "github.com/golerplate/contracts/clients/ntfr-email-sdr/v1"
	"github.com/golerplate/contracts/generated/services/ntfr/email/sdr/v1/sdrv1connect"
	grpc_interceptors "github.com/golerplate/pkg/grpc/interceptors"
)

type NotifierEmailSdrClient struct {
	client sdrv1connect.NotifierEmailSdrClient
}

func NewNotifierEmailSdrClient(ctx context.Context, notifierEmailSdrClientURL string) notifier_email_sdr_v1.NotifierEmailSdr {
	iceptorsChain := grpc_interceptors.ClientDefaultChain()

	httpClient := &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:        0,
			MaxConnsPerHost:     0,
			MaxIdleConnsPerHost: 0,
		},
		Timeout: 5 * time.Second,
	}

	return &NotifierEmailSdrClient{
		client: sdrv1connect.NewNotifierEmailSdrClient(httpClient, notifierEmailSdrClientURL, connect.WithInterceptors(iceptorsChain...)),
	}
}
