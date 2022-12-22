package create

type InputCreateUserCredentialsDto struct {
	UserId   string `json:"-"`
	Password string `json:"password"`
}

type OutputCreateUserCredentialsDto struct {
	CreatedAt string `json:"created_at"`
}
