package user_store_svc_grpc_v1

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/golerplate/pkg/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"

	pb "github.com/golerplate/contracts/generated/services/user/store/svc/v1"
)

func (c *UserStoreSvcClient) ChangePassword(ctx context.Context, userID, oldPassword, newPassword string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	req := connect.NewRequest(&pb.ChangePasswordRequest{
		UserId:      &wrapperspb.StringValue{Value: userID},
		OldPassword: &wrapperspb.StringValue{Value: oldPassword},
		NewPassword: &wrapperspb.StringValue{Value: newPassword},
	})

	_, err := c.client.ChangePassword(ctx, req)
	if err != nil {
		return grpc.TranslateFromGRPCError(ctx, err)
	}

	return nil
}
