package ntfr_email_sdr_v1

import (
	"context"
)

//go:generate mockgen -source interface.go -destination mocks/mock_ntfr_email_sdr.go -package ntfr_email_sdr_v1_mocks
type NtfrEmailSdr interface {
	SendVerificationEmail(ctx context.Context, email, username, token string) error
	SendResetPasswordEmail(ctx context.Context, email, username, token string) error
}
