package ntfr_email_sdr_grpc_v1

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/golerplate/pkg/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"

	pb "github.com/golerplate/contracts/generated/services/ntfr/email/sdr/v1"
)

func (c *NtfrEmailSdrClient) SendResetPasswordEmail(ctx context.Context, email string, username string, token string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	req := connect.NewRequest(&pb.SendResetPasswordEmailRequest{
		Email:    &wrapperspb.StringValue{Value: email},
		Username: &wrapperspb.StringValue{Value: username},
		Token:    &wrapperspb.StringValue{Value: token},
	})

	_, err := c.client.SendResetPasswordEmail(ctx, req)
	if err != nil {
		return grpc.TranslateFromGRPCError(ctx, err)
	}

	return nil
}
