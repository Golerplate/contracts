package models_http_common_account_v1

import "time"

type Account struct {
	User *User `json:"user"`
}

type User struct {
	ID             string    `json:"id"`
	Username       string    `json:"username"`
	IsVerified     bool      `json:"is_verified"`
	ProfilePicture string    `json:"profile_picture"`
	CreatedAt      time.Time `json:"created_at"`
}