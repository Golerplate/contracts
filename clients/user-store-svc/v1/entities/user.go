package user_store_svc_v1_entities

import "time"

type User struct {
	ID         string    `json:"id"`
	ExternalID string    `json:"external_id"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type UserCreate struct {
	ExternalID string `json:"external_id"`
	Email      string `json:"email"`
}
