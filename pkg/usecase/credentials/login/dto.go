package login

type InputUserLoginDto struct {
	Email    string
	Password string
}

type OutputUserLoginDto struct {
	AccessToken string `json:"access_token"`
}
