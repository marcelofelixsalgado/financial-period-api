package login

type InputUserLoginDto struct {
	Email    string
	Password string
}

type OutputUserLoginDto struct {
	User        userDto `json:"user"`
	AccessToken string  `json:"access_token"`
}

type userDto struct {
	Id string `json:"id"`
}
