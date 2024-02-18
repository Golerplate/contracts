package user_store_svc_grpc_v1

import (
	"context"

	user_store_svc_v1_entities "github.com/Golerplate/contracts/clients/user-store-svc/v1/entities"
	"github.com/Golerplate/pkg/grpc"
	"github.com/bufbuild/connect-go"
	"google.golang.org/protobuf/types/known/wrapperspb"

	pb "github.com/Golerplate/contracts/generated/services/user/store/svc/v1"
)

func (c *UserStoreSvcClient) GetUserByEmail(ctx context.Context, email string) (*user_store_svc_v1_entities.User, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	req := connect.NewRequest(&pb.GetUserByEmailRequest{
		Email: &wrapperspb.StringValue{Value: email},
	})

	resp, err := c.client.GetUserByEmail(ctx, req)
	if err != nil {
		return nil, grpc.TranslateFromGRPCError(ctx, err)
	}

	return &user_store_svc_v1_entities.User{
		ID:        resp.Msg.GetUser().Id.GetValue(),
		Username:  resp.Msg.GetUser().Username.GetValue(),
		Email:     resp.Msg.GetUser().Email.GetValue(),
		IsAdmin:   resp.Msg.GetUser().IsAdmin.GetValue(),
		IsBanned:  resp.Msg.GetUser().IsBanned.GetValue(),
		CreatedAt: resp.Msg.GetUser().CreatedAt.AsTime(),
		UpdatedAt: resp.Msg.GetUser().UpdatedAt.AsTime(),
	}, nil
}
