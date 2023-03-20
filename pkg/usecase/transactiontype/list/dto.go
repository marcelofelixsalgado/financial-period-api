package list

type InputListGroupDto struct {
	TenantId string
}

type OutputListGroupDto struct {
	Groups []Group `json:"-"`
}

type Group struct {
	Id   string    `json:"id"`
	Code string    `json:"code"`
	Name string    `json:"name"`
	Type GroupType `json:"type"`
}

type GroupType struct {
	Code string `json:"code"`
}
