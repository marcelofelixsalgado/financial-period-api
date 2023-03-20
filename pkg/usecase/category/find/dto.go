package find

type InputFindCategoryDto struct {
	Id string
}

type OutputFindCategoryDto struct {
	Id              string                `json:"id"`
	Code            string                `json:"code"`
	Name            string                `json:"name"`
	TransactionType TransactionTypeOutput `json:"transaction_type"`
	CreatedAt       string                `json:"created_at"`
	UpdatedAt       string                `json:"updated_at"`
}

type TransactionTypeOutput struct {
	Code string `json:"code"`
	Name string `json:"name"`
}
