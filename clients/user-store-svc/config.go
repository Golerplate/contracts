package user_store_svc

type UserStoreSvcConfig struct {
	Addr string `envconfig:"USER_STORE_SVC_ADDR"`
}
