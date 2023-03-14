package update

type InputUpdateGroupDto struct {
	Id       string
	TenantId string
	Code     string    `json:"code"`
	Name     string    `json:"name"`
	Type     GroupType `json:"type"`
}

type OutputUpdateGroupDto struct {
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
