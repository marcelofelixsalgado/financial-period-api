package find

type InputFindSubCategoryDto struct {
	Id string
}

type OutputFindSubCategoryDto struct {
	Id        string   `json:"id"`
	Code      string   `json:"code"`
	Name      string   `json:"name"`
	Category  Category `json:"category"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}

type Category struct {
	Id              string          `json:"id"`
	Code            string          `json:"code"`
	Name            string          `json:"name"`
	TransactionType TransactionType `json:"transaction_type"`
}

type TransactionType struct {
	Code string `json:"code"`
	Name string `json:"name"`
}
