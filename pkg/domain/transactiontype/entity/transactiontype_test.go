package entity_test

import (
	"testing"

	"github.com/marcelofelixsalgado/financial-period-api/pkg/domain/transactiontype/entity"
)

type testCase struct {
	code     string
	name     string
	expected string
}

func TestNewTransactionTypeSuccess(t *testing.T) {

	testCases := []testCase{
		{
			code: "EXP",
			name: "Expense",
		},
	}

	for _, testCase := range testCases {
		received, err := entity.NewTransactionType(testCase.code, testCase.name)
		if err != nil {
			t.Errorf("Should not return an error: %s", err)
		}
		if testCase.code != received.GetCode() {
			t.Errorf("Code expected: %s - received: %s", testCase.code, received.GetCode())
		}
		if testCase.name != received.GetName() {
			t.Errorf("Name expected: %s - received: %s", testCase.name, received.GetName())
		}
	}
}
