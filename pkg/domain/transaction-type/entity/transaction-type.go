package entity

import (
	"errors"
	"strings"
)

type ITransactionType interface {
	GetCode() string
	GetName() string
}

type TransactionType struct {
	code string
	name string
}

func NewTransactionType(code string, name string) (ITransactionType, error) {
	transactionType := TransactionType{
		code: code,
		name: name,
	}
	transactionType.format()
	if err := transactionType.validate(); err != nil {
		return nil, err
	}

	return transactionType, nil
}

func (transactionType TransactionType) GetCode() string {
	return transactionType.code
}

func (transactionType TransactionType) GetName() string {
	return transactionType.name
}

func (transactionType TransactionType) format() {
	transactionType.code = strings.TrimSpace(transactionType.code)
	transactionType.name = strings.TrimSpace(transactionType.name)
}

func (transactionType *TransactionType) validate() error {

	if transactionType.code == "" {
		return errors.New("code is required")
	}

	if transactionType.name == "" {
		return errors.New("name is required")
	}

	return nil
}
