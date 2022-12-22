package update

type InputUpdateUserDto struct {
	Id    string `json:"-"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

type OutputUpdateUserDto struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
