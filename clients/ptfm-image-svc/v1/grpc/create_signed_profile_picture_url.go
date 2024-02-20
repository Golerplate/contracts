package ptfm_image_svc_grpc_v1

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/golerplate/pkg/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"

	pb "github.com/golerplate/contracts/generated/services/ptfm/image/svc/v1"
)

func (c *PtfmImageSvcClient) CreateSignedProfilePictureUrl(ctx context.Context, userID string) (string, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	req := connect.NewRequest(&pb.CreateSignedProfilePictureUrlRequest{
		UserId: wrapperspb.String(userID),
	})

	resp, err := c.client.CreateSignedProfilePictureUrl(ctx, req)
	if err != nil {
		return "", grpc.TranslateFromGRPCError(ctx, err)
	}

	return resp.Msg.GetProfilePictureUrl().GetValue(), nil
}
