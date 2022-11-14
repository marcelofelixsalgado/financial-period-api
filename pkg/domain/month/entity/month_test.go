package entity_test

import (
	"marcelofelixsalgado/financial-month-api/pkg/domain/month/entity"
	"testing"
	"time"
)

type testCase struct {
	code      string
	name      string
	year      int
	startDate time.Time
	endDate   time.Time
	expected  string
}

func TestNewMonthSuccess(t *testing.T) {

	testCases := []testCase{
		{
			code:      "11",
			name:      "November",
			year:      2022,
			startDate: time.Now(),
			endDate:   time.Now(),
		},
	}

	for _, testCase := range testCases {
		received, err := entity.NewMonth(testCase.code, testCase.name, testCase.year, testCase.startDate, testCase.endDate)
		if err != nil {
			t.Errorf("Should not return an error: %s", err)
		}
		if testCase.code != received.GetCode() {
			t.Errorf("Code expected: %s - received: %s", testCase.code, received.GetCode())
		}
		if testCase.name != received.GetName() {
			t.Errorf("Name expected: %s - received: %s", testCase.name, received.GetName())
		}
		if testCase.year != received.GetYear() {
			t.Errorf("Year expected: %d - received: %d", testCase.year, received.GetYear())
		}
		if testCase.startDate != received.GetStartDate() {
			t.Errorf("StartDate expected: %v - received: %v", testCase.startDate, received.GetStartDate())
		}
		if testCase.endDate != received.GetEndDate() {
			t.Errorf("EndDate expected: %v - received: %v", testCase.endDate, received.GetEndDate())
		}
	}
}

func TestNewMonthInvalidCode(t *testing.T) {
	testCase := testCase{
		code:      "",
		name:      "November",
		year:      2022,
		startDate: time.Now(),
		endDate:   time.Now(),
		expected:  "code is required",
	}
	_, err := entity.NewMonth(testCase.code, testCase.name, testCase.year, testCase.startDate, testCase.endDate)
	if err.Error() != testCase.expected {
		t.Errorf("Error expected: %s - Error received: %s", testCase.expected, err)
	}
}

func TestNewMonthInvalidName(t *testing.T) {
	testCase := testCase{
		code:      "11",
		name:      "",
		year:      2022,
		startDate: time.Now(),
		endDate:   time.Now(),
		expected:  "name is required",
	}
	_, err := entity.NewMonth(testCase.code, testCase.name, testCase.year, testCase.startDate, testCase.endDate)

	if err.Error() != testCase.expected {
		t.Errorf("Error expected: %s - Error received: %s", testCase.expected, err)
	}
}

func TestNewMonthInvalidYear(t *testing.T) {
	testCase := testCase{
		code:      "11",
		name:      "Novembro",
		year:      0,
		startDate: time.Now(),
		endDate:   time.Now(),
		expected:  "year is required",
	}
	_, err := entity.NewMonth(testCase.code, testCase.name, testCase.year, testCase.startDate, testCase.endDate)
	if err.Error() != testCase.expected {
		t.Errorf("Error expected: %s - Error received: %s", testCase.expected, err)
	}
}

func TestNewMonthInvalidStartDate(t *testing.T) {
	testCase := testCase{
		code:     "11",
		name:     "November",
		year:     2022,
		endDate:  time.Now(),
		expected: "start date is required",
	}
	_, err := entity.NewMonth(testCase.code, testCase.name, testCase.year, testCase.startDate, testCase.endDate)
	if err.Error() != testCase.expected {
		t.Errorf("Error expected: %s - Error received: %s", testCase.expected, err)
	}
}

func TestNewMonthInvalidEndDate(t *testing.T) {
	testCase := testCase{
		code:      "11",
		name:      "November",
		year:      2022,
		startDate: time.Now(),
		expected:  "end date is required",
	}
	_, err := entity.NewMonth(testCase.code, testCase.name, testCase.year, testCase.startDate, testCase.endDate)
	if err.Error() != testCase.expected {
		t.Errorf("Error expected: %s - Error received: %s", testCase.expected, err)
	}
}
