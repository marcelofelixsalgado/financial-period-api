package create

type InputCreateUserDto struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

type OutputCreateUserDto struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}
