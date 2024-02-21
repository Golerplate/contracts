package user_store_svc_grpc_v1

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/golerplate/pkg/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"

	pb "github.com/golerplate/contracts/generated/services/user/store/svc/v1"
)

func (c *UserStoreSvcClient) VerifyPassword(ctx context.Context, userID, password string) (bool, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	req := connect.NewRequest(&pb.VerifyPasswordRequest{
		UserId:   &wrapperspb.StringValue{Value: userID},
		Password: &wrapperspb.StringValue{Value: password},
	})

	resp, err := c.client.VerifyPassword(ctx, req)
	if err != nil {
		return false, grpc.TranslateFromGRPCError(ctx, err)
	}

	return resp.Msg.IsValid.GetValue(), nil
}
