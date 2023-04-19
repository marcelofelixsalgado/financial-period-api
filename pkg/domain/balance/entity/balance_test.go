package entity_test

import (
	"fmt"
	"testing"

	"github.com/marcelofelixsalgado/financial-period-api/pkg/domain/balance/entity"
)

type testCase struct {
	tenantId     string
	periodId     string
	categoryId   string
	actualAmount float32
	limitAmount  float32
	expected     string
}

func TestNewBalanceSuccess(t *testing.T) {
	testCase := testCase{
		tenantId:     "456",
		periodId:     "789",
		categoryId:   "clothes",
		actualAmount: 1000,
		limitAmount:  1500,
	}

	balance, err := entity.Create(testCase.tenantId, testCase.periodId, testCase.categoryId, testCase.actualAmount, testCase.limitAmount)
	if err != err {
		t.Errorf("Should not return an error: %s", err)
	}
	if testCase.tenantId != balance.GetTenantId() {
		t.Errorf("Tenant Id expected %s - received: %s", testCase.tenantId, balance.GetTenantId())
	}
	if testCase.periodId != balance.GetPeriodId() {
		t.Errorf("Id expected %s - received: %s", testCase.periodId, balance.GetPeriodId())
	}
	if testCase.categoryId != balance.GetCategoryId() {
		t.Errorf("Id expected %s - received: %s", testCase.categoryId, balance.GetCategoryId())
	}
	if testCase.actualAmount != balance.GetActualAmount() {
		t.Errorf("Id expected %f - received: %f", testCase.actualAmount, balance.GetActualAmount())
	}
	if testCase.limitAmount != balance.GetLimitAmount() {
		t.Errorf("Id expected %f - received: %f", testCase.limitAmount, balance.GetLimitAmount())
	}
}

func TestInvalidData(t *testing.T) {
	testCases := []testCase{
		{
			tenantId:     "",
			periodId:     "789",
			categoryId:   "clothes",
			actualAmount: 1000,
			limitAmount:  1500,
			expected:     "tenant id is required",
		},
		{
			tenantId:     "456",
			periodId:     "",
			categoryId:   "clothes",
			actualAmount: 1000,
			limitAmount:  1500,
			expected:     "period is required",
		},
		{
			tenantId:     "456",
			periodId:     "789",
			categoryId:   "",
			actualAmount: 1000,
			limitAmount:  1500,
			expected:     "category is required",
		},
	}

	for _, testCase := range testCases {
		_, err := entity.Create(testCase.tenantId, testCase.periodId, testCase.categoryId, testCase.actualAmount, testCase.limitAmount)
		if err == nil || (err.Error() != testCase.expected) {
			t.Errorf(formatErrorDiff(testCase.expected, err))
		}
	}
}

func formatErrorDiff(expected string, received error) string {
	return fmt.Sprintf("Error expected: %s - Error received: %s", expected, received)
}
