package update

type InputUpdateBalanceDto struct {
	Id           string
	TenantId     string
	ActualAmount float32 `json:"actual_amount"`
}

type OutputUpdateBalanceDto struct {
	Id           string  `json:"id"`
	PeriodId     string  `json:"period_id"`
	CategoryId   string  `json:"category_id"`
	ActualAmount float32 `json:"actual_amount"`
	LimitAmount  float32 `json:"limit_amout"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
}
