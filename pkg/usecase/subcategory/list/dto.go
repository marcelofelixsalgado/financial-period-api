package list

type InputListSubCategoryDto struct {
	TenantId string
}

type OutputListSubCategoryDto struct {
	SubCategories []SubCategory `json:"-"`
}

type SubCategory struct {
	Id       string   `json:"id"`
	Code     string   `json:"code"`
	Name     string   `json:"name"`
	Category Category `json:"category"`
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
