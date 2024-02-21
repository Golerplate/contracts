package models_http_common_account_v1

type PrivateAccount struct {
	User *PrivateUser `json:"user"`
}

type PrivateUser struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	IsAdmin  bool   `json:"is_admin"`
}
