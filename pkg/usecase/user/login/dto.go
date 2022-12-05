package login

type InputUserLoginDto struct {
	Email    string
	Password string
}

type OutputUserLoginDto struct {
	Email       string `json:"email"`
	AccessToken string `json:"access_token"`
}
