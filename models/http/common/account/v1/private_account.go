package models_http_common_account_v1

type PrivateAccount struct {
	User *PrivateUser `json:"user"`
}

type PrivateUser struct {
	ID             string `json:"id"`
	Username       string `json:"username"`
	Email          string `json:"email"`
	IsVerified     bool   `json:"is_verified"`
	ProfilePicture string `json:"profile_picture"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}
