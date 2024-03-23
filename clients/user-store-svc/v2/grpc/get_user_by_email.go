package user_store_svc_grpc_v2

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/golerplate/pkg/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"

	user_store_svc_v2_entities "github.com/golerplate/contracts/clients/user-store-svc/v2/entities"
	pb "github.com/golerplate/contracts/generated/services/user/store/svc/v2"
)

func (c *UserStoreSvcClient) GetUserByEmail(ctx context.Context, email string) (*user_store_svc_v2_entities.User, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	req := connect.NewRequest(&pb.GetUserByEmailRequest{
		Email: &wrapperspb.StringValue{Value: email},
	})

	resp, err := c.client.GetUserByEmail(ctx, req)
	if err != nil {
		return nil, grpc.TranslateFromGRPCError(ctx, err)
	}

	return &user_store_svc_v2_entities.User{
		ID:        resp.Msg.GetUser().Id.GetValue(),
		Username:  resp.Msg.GetUser().Username.GetValue(),
		Email:     resp.Msg.GetUser().Email.GetValue(),
		IsBanned:  resp.Msg.GetUser().IsBanned.GetValue(),
		CreatedAt: resp.Msg.GetUser().CreatedAt.AsTime(),
		UpdatedAt: resp.Msg.GetUser().UpdatedAt.AsTime(),
	}, nil
}
