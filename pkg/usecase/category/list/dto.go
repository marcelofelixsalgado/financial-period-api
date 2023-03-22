package list

type InputListCategoryDto struct {
	TenantId string
}

type OutputListCategoryDto struct {
	Categories []Category `json:"-"`
}

type Category struct {
	Id              string                `json:"id"`
	Code            string                `json:"code"`
	Name            string                `json:"name"`
	TransactionType TransactionTypeOutput `json:"transaction_type"`
}

type TransactionTypeOutput struct {
	Code string `json:"code"`
	Name string `json:"name"`
}
