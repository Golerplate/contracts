package user_store_svc_grpc_v1

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/golerplate/pkg/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"

	user_store_svc_v1_entities "github.com/golerplate/contracts/clients/user-store-svc/v1/entities"
	pb "github.com/golerplate/contracts/generated/services/user/store/svc/v1"
)

func (c *UserStoreSvcClient) CreateUser(ctx context.Context, user *user_store_svc_v1_entities.UserCreate) (*user_store_svc_v1_entities.User, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	req := connect.NewRequest(&pb.CreateUserRequest{
		Username: &wrapperspb.StringValue{Value: user.Username},
		Email:    &wrapperspb.StringValue{Value: user.Email},
		Password: &wrapperspb.StringValue{Value: user.Password},
	})

	resp, err := c.client.CreateUser(ctx, req)
	if err != nil {
		return nil, grpc.TranslateFromGRPCError(ctx, err)
	}

	return &user_store_svc_v1_entities.User{
		ID:             resp.Msg.GetUser().Id.GetValue(),
		Username:       resp.Msg.GetUser().Username.GetValue(),
		Email:          resp.Msg.GetUser().Email.GetValue(),
		IsVerified:     resp.Msg.GetUser().IsVerified.GetValue(),
		ProfilePicture: resp.Msg.GetUser().ProfilePicture.GetValue(),
		CreatedAt:      resp.Msg.GetUser().CreatedAt.AsTime(),
		UpdatedAt:      resp.Msg.GetUser().UpdatedAt.AsTime(),
	}, nil
}
