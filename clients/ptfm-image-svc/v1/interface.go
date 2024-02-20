package ptfm_image_svc_v1

import (
	"context"
)

//go:generate mockgen -source interface.go -destination mocks/mock_ptfm_image_svc.go -package ptfm_image_svc_v1_mocks
type PtfmImageSvc interface {
	CreateSignedProfilePictureUrl(ctx context.Context, userID string) (string, error)
}
