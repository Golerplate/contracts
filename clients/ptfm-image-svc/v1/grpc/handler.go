package ptfm_image_svc_grpc_v1

import (
	"context"
	"net/http"
	"time"

	"github.com/bufbuild/connect-go"
	grpc_interceptors "github.com/golerplate/pkg/grpc/interceptors"

	ptfm_image_svc_v1 "github.com/golerplate/contracts/clients/ptfm-image-svc/v1"
	"github.com/golerplate/contracts/generated/services/ptfm/image/svc/v1/svcv1connect"
)

type PtfmImageSvcClient struct {
	client svcv1connect.PtfmImageSvcClient
}

func NewPtfmImageSvcClient(ctx context.Context, ptfmImageSvcClientURL string) ptfm_image_svc_v1.PtfmImageSvc {
	iceptorsChain := grpc_interceptors.ClientDefaultChain()

	httpClient := &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:        0,
			MaxConnsPerHost:     0,
			MaxIdleConnsPerHost: 0,
		},
		Timeout: 5 * time.Second,
	}

	return &PtfmImageSvcClient{
		client: svcv1connect.NewPtfmImageSvcClient(httpClient, ptfmImageSvcClientURL, connect.WithInterceptors(iceptorsChain...)),
	}
}
