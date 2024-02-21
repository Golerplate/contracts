package models_http_common_account_v1

import "time"

type PrivateAccount struct {
	User *PrivateUser `json:"user"`
}

type PrivateUser struct {
	ID               string    `json:"id"`
	Username         string    `json:"username"`
	Email            string    `json:"email"`
	IsAdmin          bool      `json:"is_admin"`
	IsBanned         bool      `json:"is_banned"`
	HasVerifiedEmail bool      `json:"has_verified_email"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
