package entity_test

import (
	"marcelofelixsalgado/financial-period-api/pkg/domain/user/entity"
	"strings"
	"testing"
)

type testCase struct {
	name     string
	password string
	phone    string
	email    string
	expected string
}

func TestNewUserSuccess(t *testing.T) {

	testCases := []testCase{
		{
			name:     "John",
			password: "123",
			phone:    "123456",
			email:    "john@test.com",
		},
	}

	for _, testCase := range testCases {
		received, err := entity.Create(testCase.name, testCase.password, testCase.phone, testCase.email)
		if err != nil {
			t.Errorf("Should not return an error: %s", err)
		}
		if testCase.name != received.GetName() {
			t.Errorf("Name expected: %s - received: %s", testCase.name, received.GetName())
		}
		if testCase.password != received.GetPassword() {
			t.Errorf("Password expected: %s - received: %s", testCase.password, received.GetPassword())
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
		name:     "     	John      ",
		password: "123",
		phone:    "123456",
		email:    "john@test.com",
	}
	expectedName := "John"

	received, err := entity.Create(testCase.name, testCase.password, testCase.phone, testCase.email)
	if err != nil {
		t.Errorf("Should not return an error: %s", err)
	}

	if strings.Compare(expectedName, received.GetName()) != 0 {
		t.Errorf("Name expected: [%s] - received: [%s]", expectedName, received.GetName())
	}
}

func TestNewUserInvalidName(t *testing.T) {
	testCase := testCase{
		name:     "",
		password: "123",
		phone:    "123456",
		email:    "john@test.com",
		expected: "name is required",
	}
	_, err := entity.Create(testCase.name, testCase.password, testCase.phone, testCase.email)
	if err == nil || (err.Error() != testCase.expected) {
		t.Errorf("Error expected: %s - Error received: %s", testCase.expected, err)
	}
}

func TestNewUserInvalidPassword(t *testing.T) {
	testCase := testCase{
		name:     "John",
		password: "",
		phone:    "123456",
		email:    "john@test.com",
		expected: "password is required",
	}
	_, err := entity.Create(testCase.name, testCase.password, testCase.phone, testCase.email)
	if err == nil || (err.Error() != testCase.expected) {
		t.Errorf("Error expected: %s - Error received: %s", testCase.expected, err)
	}
}

func TestNewUserInvalidPhone(t *testing.T) {
	testCase := testCase{
		name:     "John",
		password: "123",
		phone:    "",
		email:    "john@test.com",
		expected: "phone is required",
	}
	_, err := entity.Create(testCase.name, testCase.password, testCase.phone, testCase.email)
	if err == nil || (err.Error() != testCase.expected) {
		t.Errorf("Error expected: %s - Error received: %s", testCase.expected, err)
	}
}

func TestNewUserInvalidEmail(t *testing.T) {
	testCase := testCase{
		name:     "John",
		password: "123",
		phone:    "123456",
		email:    "",
		expected: "email is required",
	}
	_, err := entity.Create(testCase.name, testCase.password, testCase.phone, testCase.email)
	if err == nil || (err.Error() != testCase.expected) {
		t.Errorf("Error expected: %s - Error received: %s", testCase.expected, err)
	}
}