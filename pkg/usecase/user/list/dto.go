package list

type InputListUserDto struct {
}

type User struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

type OutputListUserDto struct {
	Users []User `json:"-"`
}
