package find

type InputFindUserDto struct {
	Id string
}

type OutputFindUserDto struct {
	Id        string    `json:"id"`
	Tenant    tenantDto `json:"tenant"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at,omitempty"`
}

type tenantDto struct {
	Id string `json:"id"`
}
