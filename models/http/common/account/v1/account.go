package models_http_common_account_v1

import "time"

type Account struct {
	User *User `json:"user"`
}

type User struct {
	ID               string    `json:"id"`
	Username         string    `json:"username"`
	IsAdmin          bool      `json:"is_admin"`
	IsBanned         bool      `json:"is_banned"`
	HasVerifiedEmail bool      `json:"has_verified_email"`
	CreatedAt        time.Time `json:"created_at"`
}
