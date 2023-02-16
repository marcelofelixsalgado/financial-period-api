package create

import "time"

type InputCreateBalanceDto struct {
	TenantId     string
	PeriodId     string  `json:"period_id"`
	CategoryId   string  `json:"category_id"`
	ActualAmount float32 `json:"actual_amount"`
	LimitAmount  float32 `json:"limit_amount"`
}

type OutputCreateBalanceDto struct {
	Id           string    `json:"id"`
	PeriodId     string    `json:"period_id"`
	CategoryId   string    `json:"category_id"`
	ActualAmount float32   `json:"actual_amount"`
	LimitAmount  float32   `json:"limit_amout"`
	CreatedAt    time.Time `json:"created_at"`
}
