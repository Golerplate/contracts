package user_gtw_v1_entities

type CreateUserRequest struct {
	ExternalID string `json:"external_id"`
	Email      string `json:"email"`
}
