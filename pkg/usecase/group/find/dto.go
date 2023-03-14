package find

type InputFindGroupDto struct {
	Id string
}

type OutputFindGroupDto struct {
	Id        string    `json:"id"`
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	Type      GroupType `json:"type"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
}

type GroupType struct {
	Code string `json:"code"`
}
