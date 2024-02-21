package sess_store_svc_grpc_v1

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/golerplate/pkg/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"

	sess_store_svc_v1_entities "github.com/golerplate/contracts/clients/sess-store-svc/v1/entities"
	pb "github.com/golerplate/contracts/generated/services/sess/store/svc/v1"
)

func (c *SessStoreSvcClient) GetSessionByID(ctx context.Context, id string) (*sess_store_svc_v1_entities.Session, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	req := connect.NewRequest(&pb.GetSessionByIDRequest{
		Id: &wrapperspb.StringValue{Value: id},
	})

	resp, err := c.client.GetSessionByID(ctx, req)
	if err != nil {
		return nil, grpc.TranslateFromGRPCError(ctx, err)
	}

	return &sess_store_svc_v1_entities.Session{
		ID:        resp.Msg.GetSession().Id.GetValue(),
		UserID:    resp.Msg.GetSession().UserId.GetValue(),
		Token:     resp.Msg.GetSession().Token.GetValue(),
		CreatedAt: resp.Msg.GetSession().CreatedAt.AsTime(),
	}, nil
}