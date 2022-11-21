package entity_test

import (
	"marcelofelixsalgado/financial-period-api/pkg/domain/period/entity"
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

func TestNewPeriodSuccess(t *testing.T) {

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
		received, err := entity.NewPeriod(testCase.code, testCase.name, testCase.year, testCase.startDate, testCase.endDate)
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
		if received.GetCreatedAt().IsZero() {
			t.Errorf("CreatedAt must not be zero")
		}
	}
}

func TestNewPeriodInvalidCode(t *testing.T) {
	testCase := testCase{
		code:      "",
		name:      "November",
		year:      2022,
		startDate: time.Now(),
		endDate:   time.Now(),
		expected:  "code is required",
	}
	_, err := entity.NewPeriod(testCase.code, testCase.name, testCase.year, testCase.startDate, testCase.endDate)
	if err.Error() != testCase.expected {
		t.Errorf("Error expected: %s - Error received: %s", testCase.expected, err)
	}
}

func TestNewPeriodInvalidName(t *testing.T) {
	testCase := testCase{
		code:      "11",
		name:      "",
		year:      2022,
		startDate: time.Now(),
		endDate:   time.Now(),
		expected:  "name is required",
	}
	_, err := entity.NewPeriod(testCase.code, testCase.name, testCase.year, testCase.startDate, testCase.endDate)

	if err.Error() != testCase.expected {
		t.Errorf("Error expected: %s - Error received: %s", testCase.expected, err)
	}
}

func TestNewPeriodInvalidYear(t *testing.T) {
	testCase := testCase{
		code:      "11",
		name:      "Novembro",
		year:      0,
		startDate: time.Now(),
		endDate:   time.Now(),
		expected:  "year is required",
	}
	_, err := entity.NewPeriod(testCase.code, testCase.name, testCase.year, testCase.startDate, testCase.endDate)
	if err.Error() != testCase.expected {
		t.Errorf("Error expected: %s - Error received: %s", testCase.expected, err)
	}
}

func TestNewPeriodInvalidStartDate(t *testing.T) {
	testCase := testCase{
		code:     "11",
		name:     "November",
		year:     2022,
		endDate:  time.Now(),
		expected: "start date is required",
	}
	_, err := entity.NewPeriod(testCase.code, testCase.name, testCase.year, testCase.startDate, testCase.endDate)
	if err.Error() != testCase.expected {
		t.Errorf("Error expected: %s - Error received: %s", testCase.expected, err)
	}
}

func TestNewPeriodInvalidEndDate(t *testing.T) {
	testCase := testCase{
		code:      "11",
		name:      "November",
		year:      2022,
		startDate: time.Now(),
		expected:  "end date is required",
	}
	_, err := entity.NewPeriod(testCase.code, testCase.name, testCase.year, testCase.startDate, testCase.endDate)
	if err.Error() != testCase.expected {
		t.Errorf("Error expected: %s - Error received: %s", testCase.expected, err)
	}
}

func TestNewPeriodEqualDates(t *testing.T) {
	sameDate := time.Now()
	testCase := testCase{
		code:      "11",
		name:      "November",
		year:      2022,
		startDate: sameDate,
		endDate:   sameDate,
		expected:  "start date must be greater than the end date",
	}
	_, err := entity.NewPeriod(testCase.code, testCase.name, testCase.year, testCase.startDate, testCase.endDate)
	if err.Error() != testCase.expected {
		t.Errorf("Error expected: %s - Error received: %s", testCase.expected, err)
	}
}

func TestNewPeriodInvalidDates(t *testing.T) {
	sameDate := time.Now()
	testCase := testCase{
		code:      "11",
		name:      "November",
		year:      2022,
		startDate: sameDate.Add(24 * time.Hour),
		endDate:   sameDate,
		expected:  "start date must be greater than the end date",
	}
	_, err := entity.NewPeriod(testCase.code, testCase.name, testCase.year, testCase.startDate, testCase.endDate)
	if err.Error() != testCase.expected {
		t.Errorf("Error expected: %s - Error received: %s", testCase.expected, err)
	}
}
