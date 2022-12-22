package update

type InputUpdateUserCredentialsDto struct {
	Id              string `json:"-"`
	UserId          string `json:"user_id"`
	CurrentPassword string `json:"current_password"`
	NewPassword     string `json:"new_password"`
}

type OutputUpdateUserCredentialsDto struct {
	Id        string `json:"id"`
	UserId    string `json:"user_id"`
	Password  string `json:"-"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
