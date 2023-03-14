package entity_test

import (
	"fmt"
	. "marcelofelixsalgado/financial-period-api/pkg/domain/group/entity"
	"strings"
	"testing"
)

type testCase struct {
	tenantId  string
	code      string
	name      string
	groupType GroupType
	expected  string
}

func TestNewGroupSuccess(t *testing.T) {

	testCases := []testCase{
		{
			tenantId: "123",
			code:     "DF",
			name:     "Despesa fixa",
			groupType: GroupType{
				Code: "EXP",
			},
		},
	}

	for _, testCase := range testCases {
		received, err := Create(testCase.tenantId, testCase.code, testCase.name, testCase.groupType)
		if err != nil {
			t.Errorf("Should not return an error: %s", err)
		}
		if testCase.code != received.GetCode() {
			t.Errorf("Code expected: %s - received: %s", testCase.code, received.GetCode())
		}
		if testCase.name != received.GetName() {
			t.Errorf("Name expected: %s - received: %s", testCase.name, received.GetName())
		}
		if testCase.groupType.Code != received.GetGroupType().Code {
			t.Errorf("Code expected: %s - received: %s", testCase.code, received.GetCode())
		}
		if received.GetCreatedAt().IsZero() {
			t.Errorf("CreatedAt must not be zero")
		}
	}
}

func TestNewPeriodTrimSpaces(t *testing.T) {
	testCase := testCase{
		tenantId: "123",
		code:     "   DF    ",
		name:     "     Despesa fixa    ",
		groupType: GroupType{
			Code: "    EXP   ",
		},
	}
	expectedCode := "DF"
	expectedName := "Despesa fixa"
	expectedType := "EXP"

	received, err := Create(testCase.tenantId, testCase.code, testCase.name, testCase.groupType)
	if err != nil {
		t.Errorf("Should not return an error: %s", err)
	}

	if strings.Compare(expectedCode, received.GetCode()) != 0 {
		t.Errorf("Code expected: [%s] - received: [%s]", expectedCode, received.GetCode())
	}
	if strings.Compare(expectedName, received.GetName()) != 0 {
		t.Errorf("Name expected: [%s] - received: [%s]", expectedName, received.GetName())
	}
	if strings.Compare(expectedType, received.GetGroupType().Code) != 0 {
		t.Errorf("Type expected: [%s] - received: [%s]", expectedType, received.GetGroupType().Code)
	}
}

func TestNewPeriodInvalidTenantId(t *testing.T) {
	testCase := testCase{
		tenantId: "",
		code:     "DF",
		name:     "Despesa fixa",
		groupType: GroupType{
			Code: "EXP",
		},
		expected: "tenant id is required",
	}
	_, err := Create(testCase.tenantId, testCase.code, testCase.name, testCase.groupType)
	if err == nil || (err.Error() != testCase.expected) {
		t.Errorf(formatErrorDiff(testCase.expected, err))
	}
}

func TestNewPeriodInvalidCode(t *testing.T) {
	testCase := testCase{
		tenantId: "123",
		code:     "",
		name:     "Despesa fixa",
		groupType: GroupType{
			Code: "EXP",
		},
		expected: "code is required",
	}
	_, err := Create(testCase.tenantId, testCase.code, testCase.name, testCase.groupType)
	if err == nil || (err.Error() != testCase.expected) {
		t.Errorf(formatErrorDiff(testCase.expected, err))
	}
}

func TestNewPeriodInvalidName(t *testing.T) {
	testCase := testCase{
		tenantId: "123",
		code:     "DF",
		name:     "",
		groupType: GroupType{
			Code: "EXP",
		},
		expected: "name is required",
	}
	_, err := Create(testCase.tenantId, testCase.code, testCase.name, testCase.groupType)
	if err == nil || (err.Error() != testCase.expected) {
		t.Errorf(formatErrorDiff(testCase.expected, err))
	}
}

func TestNewPeriodInvalidGroupType(t *testing.T) {
	testCase := testCase{
		tenantId: "123",
		code:     "DF",
		name:     "Despesa fixa",
		groupType: GroupType{
			Code: "INV",
		},
		expected: "invalid group type",
	}
	_, err := Create(testCase.tenantId, testCase.code, testCase.name, testCase.groupType)
	if err == nil || (err.Error() != testCase.expected) {
		t.Errorf(formatErrorDiff(testCase.expected, err))
	}
}

func formatErrorDiff(expected string, received error) string {
	return fmt.Sprintf("Error expected: %s - Error received: %s", expected, received)
}
