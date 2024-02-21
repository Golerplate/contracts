package user_store_svc_grpc_v1

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/golerplate/pkg/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"

	user_store_svc_v1_entities "github.com/golerplate/contracts/clients/user-store-svc/v1/entities"
	pb "github.com/golerplate/contracts/generated/services/user/store/svc/v1"
)

func (c *UserStoreSvcClient) GetUserByID(ctx context.Context, id string) (*user_store_svc_v1_entities.User, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	req := connect.NewRequest(&pb.GetUserByIDRequest{
		Id: &wrapperspb.StringValue{Value: id},
	})

	resp, err := c.client.GetUserByID(ctx, req)
	if err != nil {
		return nil, grpc.TranslateFromGRPCError(ctx, err)
	}

	return &user_store_svc_v1_entities.User{
		ID:               resp.Msg.GetUser().Id.GetValue(),
		Username:         resp.Msg.GetUser().Username.GetValue(),
		Email:            resp.Msg.GetUser().Email.GetValue(),
		IsAdmin:          resp.Msg.GetUser().IsAdmin.GetValue(),
		IsBanned:         resp.Msg.GetUser().IsBanned.GetValue(),
		HasVerifiedEmail: resp.Msg.GetUser().HasVerifiedEmail.GetValue(),
		CreatedAt:        resp.Msg.GetUser().CreatedAt.AsTime(),
		UpdatedAt:        resp.Msg.GetUser().UpdatedAt.AsTime(),
	}, nil
}
