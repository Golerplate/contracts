package models_http_common_account_v1

type Account struct {
	User *User `json:"user"`
}

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}
