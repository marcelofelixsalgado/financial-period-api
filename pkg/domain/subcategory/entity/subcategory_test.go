package entity_test

import (
	"fmt"
	"strings"
	"testing"

	category "github.com/marcelofelixsalgado/financial-period-api/pkg/domain/category/entity"
	. "github.com/marcelofelixsalgado/financial-period-api/pkg/domain/subcategory/entity"
	transactionType "github.com/marcelofelixsalgado/financial-period-api/pkg/domain/transactiontype/entity"
)

type testCase struct {
	tenantId string
	code     string
	name     string
	category category.ICategory
	expected string
}

func TestNewSubCategorySuccess(t *testing.T) {

	tenantId := "123"
	transactionType, _ := transactionType.NewTransactionType("EXP", "Expense")
	category, _ := category.Create(tenantId, "DV", "Despesa variável", transactionType)
	testCases := []testCase{
		{
			tenantId: tenantId,
			code:     "TR",
			name:     "Transporte",
			category: category,
		},
	}

	for _, testCase := range testCases {
		received, err := Create(testCase.tenantId, testCase.code, testCase.name, testCase.category)
		if err != nil {
			t.Errorf("Should not return an error: %s", err)
		}
		if testCase.code != received.GetCode() {
			t.Errorf("Code expected: %s - received: %s", testCase.code, received.GetCode())
		}
		if testCase.name != received.GetName() {
			t.Errorf("Name expected: %s - received: %s", testCase.name, received.GetName())
		}
		if testCase.category.GetCode() != received.GetCategory().GetCode() {
			t.Errorf("Code expected: %s - received: %s", testCase.code, received.GetCode())
		}
		if received.GetCreatedAt().IsZero() {
			t.Errorf("CreatedAt must not be zero")
		}
	}
}

func TestNewSubCategoryTrimSpaces(t *testing.T) {
	tenantId := "123"
	transactionType, _ := transactionType.NewTransactionType("EXP", "Expense")
	category, _ := category.Create(tenantId, "DV", "Despesa variável", transactionType)

	testCase := testCase{
		tenantId: tenantId,
		code:     "   TR    ",
		name:     "     Transporte    ",
		category: category,
	}
	expectedCode := "TR"
	expectedName := "Transporte"

	received, err := Create(testCase.tenantId, testCase.code, testCase.name, testCase.category)
	if err != nil {
		t.Errorf("Should not return an error: %s", err)
	}

	if strings.Compare(expectedCode, received.GetCode()) != 0 {
		t.Errorf("Code expected: [%s] - received: [%s]", expectedCode, received.GetCode())
	}
	if strings.Compare(expectedName, received.GetName()) != 0 {
		t.Errorf("Name expected: [%s] - received: [%s]", expectedName, received.GetName())
	}
}

func TestNewSubCategoryInvalidTenantId(t *testing.T) {
	tenantId := "123"
	transactionType, _ := transactionType.NewTransactionType("EXP", "Expense")
	category, _ := category.Create(tenantId, "DV", "Despesa variável", transactionType)

	testCase := testCase{
		tenantId: "",
		code:     "DF",
		name:     "Despesa fixa",
		category: category,
		expected: "tenant id is required",
	}
	_, err := Create(testCase.tenantId, testCase.code, testCase.name, testCase.category)
	if err == nil || (err.Error() != testCase.expected) {
		t.Errorf(formatErrorDiff(testCase.expected, err))
	}
}

func TestNewSubCategoryInvalidCode(t *testing.T) {
	tenantId := "123"
	transactionType, _ := transactionType.NewTransactionType("EXP", "Expense")
	category, _ := category.Create(tenantId, "DV", "Despesa variável", transactionType)

	testCase := testCase{
		tenantId: tenantId,
		code:     "",
		name:     "Despesa fixa",
		category: category,
		expected: "code is required",
	}
	_, err := Create(testCase.tenantId, testCase.code, testCase.name, testCase.category)
	if err == nil || (err.Error() != testCase.expected) {
		t.Errorf(formatErrorDiff(testCase.expected, err))
	}
}

func TestNewSubCategoryInvalidName(t *testing.T) {
	tenantId := "123"
	transactionType, _ := transactionType.NewTransactionType("EXP", "Expense")
	category, _ := category.Create(tenantId, "DV", "Despesa variável", transactionType)

	testCase := testCase{
		tenantId: tenantId,
		code:     "DF",
		name:     "",
		category: category,
		expected: "name is required",
	}
	_, err := Create(testCase.tenantId, testCase.code, testCase.name, testCase.category)
	if err == nil || (err.Error() != testCase.expected) {
		t.Errorf(formatErrorDiff(testCase.expected, err))
	}
}

func TestNewSubCategoryInvalidCategory(t *testing.T) {

	testCase := testCase{
		tenantId: "123",
		code:     "DF",
		name:     "Despesa fixa",
		expected: "category is required",
	}
	_, err := Create(testCase.tenantId, testCase.code, testCase.name, testCase.category)
	if err == nil || (err.Error() != testCase.expected) {
		t.Errorf(formatErrorDiff(testCase.expected, err))
	}
}

func formatErrorDiff(expected string, received error) string {
	return fmt.Sprintf("Error expected: %s - Error received: %s", expected, received)
}
