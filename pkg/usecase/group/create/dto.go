package create

type InputCreateGroupDto struct {
	TenantId string
	Code     string    `json:"code"`
	Name     string    `json:"name"`
	Type     GroupType `json:"type"`
}

type OutputCreateGroupDto struct {
	Id        string    `json:"id"`
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	Type      GroupType `json:"type"`
	CreatedAt string    `json:"created_at"`
}

type GroupType struct {
	Code string `json:"code"`
}
