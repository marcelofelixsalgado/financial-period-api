package entity_test

import (
	"fmt"
	"strings"
	"testing"

	. "github.com/marcelofelixsalgado/financial-period-api/pkg/domain/category/entity"
	transactionType "github.com/marcelofelixsalgado/financial-period-api/pkg/domain/transactiontype/entity"
)

type testCase struct {
	tenantId        string
	code            string
	name            string
	transactionType transactionType.ITransactionType
	expected        string
}

func TestNewCategorySuccess(t *testing.T) {

	transactionType, _ := transactionType.NewTransactionType("EXP", "Expense")
	testCases := []testCase{
		{
			tenantId:        "123",
			code:            "DF",
			name:            "Despesa fixa",
			transactionType: transactionType,
		},
	}

	for _, testCase := range testCases {
		received, err := Create(testCase.tenantId, testCase.code, testCase.name, testCase.transactionType)
		if err != nil {
			t.Errorf("Should not return an error: %s", err)
		}
		if testCase.code != received.GetCode() {
			t.Errorf("Code expected: %s - received: %s", testCase.code, received.GetCode())
		}
		if testCase.name != received.GetName() {
			t.Errorf("Name expected: %s - received: %s", testCase.name, received.GetName())
		}
		if testCase.transactionType.GetCode() != received.GetTransactionType().GetCode() {
			t.Errorf("Code expected: %s - received: %s", testCase.code, received.GetCode())
		}
		if received.GetCreatedAt().IsZero() {
			t.Errorf("CreatedAt must not be zero")
		}
	}
}

func TestNewCategoryTrimSpaces(t *testing.T) {
	transactionType, _ := transactionType.NewTransactionType("EXP", "Expense")
	testCase := testCase{
		tenantId:        "123",
		code:            "   DF    ",
		name:            "     Despesa fixa    ",
		transactionType: transactionType,
	}
	expectedCode := "DF"
	expectedName := "Despesa fixa"
	expectedTransactionType := "EXP"

	received, err := Create(testCase.tenantId, testCase.code, testCase.name, testCase.transactionType)
	if err != nil {
		t.Errorf("Should not return an error: %s", err)
	}

	if strings.Compare(expectedCode, received.GetCode()) != 0 {
		t.Errorf("Code expected: [%s] - received: [%s]", expectedCode, received.GetCode())
	}
	if strings.Compare(expectedName, received.GetName()) != 0 {
		t.Errorf("Name expected: [%s] - received: [%s]", expectedName, received.GetName())
	}
	if strings.Compare(expectedTransactionType, received.GetTransactionType().GetCode()) != 0 {
		t.Errorf("Type expected: [%s] - received: [%s]", expectedTransactionType, received.GetTransactionType().GetCode())
	}
}

func TestNewCategoryInvalidTenantId(t *testing.T) {
	transactionType, _ := transactionType.NewTransactionType("EXP", "Expense")
	testCase := testCase{
		tenantId:        "",
		code:            "DF",
		name:            "Despesa fixa",
		transactionType: transactionType,
		expected:        "tenant id is required",
	}
	_, err := Create(testCase.tenantId, testCase.code, testCase.name, testCase.transactionType)
	if err == nil || (err.Error() != testCase.expected) {
		t.Errorf(formatErrorDiff(testCase.expected, err))
	}
}

func TestNewCategoryInvalidCode(t *testing.T) {
	transactionType, _ := transactionType.NewTransactionType("EXP", "Expense")
	testCase := testCase{
		tenantId:        "123",
		code:            "",
		name:            "Despesa fixa",
		transactionType: transactionType,
		expected:        "code is required",
	}
	_, err := Create(testCase.tenantId, testCase.code, testCase.name, testCase.transactionType)
	if err == nil || (err.Error() != testCase.expected) {
		t.Errorf(formatErrorDiff(testCase.expected, err))
	}
}

func TestNewCategoryInvalidName(t *testing.T) {
	transactionType, _ := transactionType.NewTransactionType("EXP", "Expense")
	testCase := testCase{
		tenantId:        "123",
		code:            "DF",
		name:            "",
		transactionType: transactionType,
		expected:        "name is required",
	}
	_, err := Create(testCase.tenantId, testCase.code, testCase.name, testCase.transactionType)
	if err == nil || (err.Error() != testCase.expected) {
		t.Errorf(formatErrorDiff(testCase.expected, err))
	}
}

func TestNewCategoryInvalidTransactionType(t *testing.T) {
	testCase := testCase{
		tenantId: "123",
		code:     "DF",
		name:     "Despesa fixa",
		expected: "transaction type is required",
	}
	_, err := Create(testCase.tenantId, testCase.code, testCase.name, testCase.transactionType)
	if err == nil || (err.Error() != testCase.expected) {
		t.Errorf(formatErrorDiff(testCase.expected, err))
	}
}

func formatErrorDiff(expected string, received error) string {
	return fmt.Sprintf("Error expected: %s - Error received: %s", expected, received)
}
