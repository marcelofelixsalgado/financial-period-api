package update

type InputUpdateUserCredentialsDto struct {
	Id       string `json:"_"`
	UserId   string `json:"user_id"`
	Password string `json:"password"`
}

type OutputUpdateUserCredentialsDto struct {
	Id        string `json:"id"`
	UserId    string `json:"user_id"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
