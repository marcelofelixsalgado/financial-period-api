package list

type InputListTransactionTypeDto struct {
}

type OutputListTransactionTypeDto struct {
	TransactionTypes []TransactionType `json:"-"`
}

type TransactionType struct {
	Code string `json:"code"`
	Name string `json:"name"`
}
