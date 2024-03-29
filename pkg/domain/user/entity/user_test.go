package entity_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/marcelofelixsalgado/financial-period-api/pkg/domain/user/entity"
)

type testCase struct {
	tenantId string
	name     string
	phone    string
	email    string
	expected string
}

const email = "john@test.com"

func TestNewUserSuccess(t *testing.T) {

	testCases := []testCase{
		{
			tenantId: "111",
			name:     "John",
			phone:    "123456",
			email:    email,
		},
	}

	for _, testCase := range testCases {
		received, err := entity.Create(testCase.tenantId, testCase.name, testCase.phone, testCase.email)
		if err != nil {
			t.Errorf("Should not return an error: %s", err)
		}
		if testCase.tenantId != received.GetTenantId() {
			t.Errorf("TenantId expected: %s - received: %s", testCase.tenantId, received.GetTenantId())
		}
		if testCase.name != received.GetName() {
			t.Errorf("Name expected: %s - received: %s", testCase.name, received.GetName())
		}
		if testCase.phone != received.GetPhone() {
			t.Errorf("Phone expected: %s - received: %s", testCase.phone, received.GetPhone())
		}
		if testCase.email != received.GetEmail() {
			t.Errorf("Email expected: %s - received: %s", testCase.email, received.GetEmail())
		}
		if received.GetCreatedAt().IsZero() {
			t.Errorf("CreatedAt must not be zero")
		}
	}
}

func TestNewUserTrimSpaces(t *testing.T) {
	testCase := testCase{
		tenantId: "111",
		name:     "     	John      ",
		phone:    "123456",
		email:    email,
	}
	expectedName := "John"

	received, err := entity.Create(testCase.tenantId, testCase.name, testCase.phone, testCase.email)
	if err != nil {
		t.Errorf("Should not return an error: %s", err)
	}

	if strings.Compare(expectedName, received.GetName()) != 0 {
		t.Errorf("Name expected: [%s] - received: [%s]", expectedName, received.GetName())
	}
}

func TestNewUserInvalidName(t *testing.T) {
	testCase := testCase{
		tenantId: "111",
		name:     "",
		phone:    "123456",
		email:    email,
		expected: "name is required",
	}
	_, err := entity.Create(testCase.tenantId, testCase.name, testCase.phone, testCase.email)
	if err == nil || (err.Error() != testCase.expected) {
		t.Errorf(formatErrorDiff(testCase.expected, err))
	}
}

func TestNewUserInvalidPhone(t *testing.T) {
	testCase := testCase{
		tenantId: "111",
		name:     "John",
		phone:    "",
		email:    email,
		expected: "phone is required",
	}
	_, err := entity.Create(testCase.tenantId, testCase.name, testCase.phone, testCase.email)
	if err == nil || (err.Error() != testCase.expected) {
		t.Errorf(formatErrorDiff(testCase.expected, err))
	}
}

func TestNewUserInvalidEmail(t *testing.T) {
	testCase := testCase{
		tenantId: "111",
		name:     "John",
		phone:    "123456",
		email:    "",
		expected: "email is required",
	}
	_, err := entity.Create(testCase.tenantId, testCase.name, testCase.phone, testCase.email)
	if err == nil || (err.Error() != testCase.expected) {
		t.Errorf(formatErrorDiff(testCase.expected, err))
	}
}

func formatErrorDiff(expected string, received error) string {
	return fmt.Sprintf("Error expected: %s - Error received: %s", expected, received)
}
