package find

type InputFindTransactionTypeDto struct {
	Code string
}

type OutputFindTransactionTypeDto struct {
	Code string `json:"code"`
	Name string `json:"name"`
}
