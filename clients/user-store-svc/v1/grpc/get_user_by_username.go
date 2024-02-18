package user_store_svc_grpc_v1

import (
	"context"

	user_store_svc_entities_v1 "github.com/Golerplate/contracts/clients/user-store-svc/v1/entities"
	"github.com/Golerplate/pkg/grpc"
	"github.com/bufbuild/connect-go"
	"google.golang.org/protobuf/types/known/wrapperspb"

	pb "github.com/Golerplate/contracts/generated/services/user/store/svc/v1"
)

func (c *UserStoreSvcClient) GetUserByUsername(ctx context.Context, username string) (*user_store_svc_entities_v1.User, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	req := connect.NewRequest(&pb.GetUserByUsernameRequest{
		Username: &wrapperspb.StringValue{Value: username},
	})

	resp, err := c.client.GetUserByUsername(ctx, req)
	if err != nil {
		return nil, grpc.TranslateFromGRPCError(ctx, err)
	}

	return &user_store_svc_entities_v1.User{
		ID:        resp.Msg.GetUser().Id.GetValue(),
		Username:  resp.Msg.GetUser().Username.GetValue(),
		Email:     resp.Msg.GetUser().Email.GetValue(),
		IsAdmin:   resp.Msg.GetUser().IsAdmin.GetValue(),
		IsBanned:  resp.Msg.GetUser().IsBanned.GetValue(),
		CreatedAt: resp.Msg.GetUser().CreatedAt.AsTime(),
		UpdatedAt: resp.Msg.GetUser().UpdatedAt.AsTime(),
	}, nil
}
