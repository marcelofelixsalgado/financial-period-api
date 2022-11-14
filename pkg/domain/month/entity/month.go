package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type IMonth interface {
	GetId() string
	GetCode() string
	GetName() string
	GetYear() int
	GetStartDate() time.Time
	GetEndDate() time.Time
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time

	SetCode(string)
	SetName(string)
	SetYear(int)
	SetStartDate(time.Time)
	SetEndDate(time.Time)
	SetUpdatedAt(time.Time)
}

type Month struct {
	id        string
	code      string
	name      string
	year      int
	startDate time.Time
	endDate   time.Time
	createdAt time.Time
	updatedAt time.Time
}

func (month Month) GetId() string {
	return month.id
}

func (month Month) GetCode() string {
	return month.code
}

func (month Month) GetName() string {
	return month.name
}

func (month Month) GetYear() int {
	return month.year
}

func (month Month) GetStartDate() time.Time {
	return month.startDate
}

func (month Month) GetEndDate() time.Time {
	return month.endDate
}

func (month Month) GetCreatedAt() time.Time {
	return month.createdAt
}

func (month Month) GetUpdatedAt() time.Time {
	return month.updatedAt
}

func (month Month) SetCode(code string) {
	month.code = code
	validate(month)
}

func (month Month) SetName(name string) {
	month.name = name
	validate(month)
}

func (month Month) SetYear(year int) {
	month.year = year
	validate(month)
}

func (month Month) SetStartDate(startDate time.Time) {
	month.startDate = startDate
	validate(month)
}

func (month Month) SetEndDate(endDate time.Time) {
	month.endDate = endDate
	validate(month)
}

func (month Month) SetUpdatedAt(updatedAt time.Time) {
	month.updatedAt = updatedAt
	validate(month)
}

func NewMonth(code string, name string, year int, startDate time.Time, endDate time.Time) (IMonth, error) {
	month := Month{
		id:        uuid.NewString(),
		code:      code,
		name:      name,
		year:      year,
		startDate: startDate,
		endDate:   endDate,
		createdAt: time.Now(),
	}
	if err := validate(month); err != nil {
		return nil, err
	}
	return month, nil
}

func validate(month Month) error {
	if month.id == "" {
		return errors.New("id is required")
	}

	if month.code == "" {
		return errors.New("code is required")
	}

	if month.name == "" {
		return errors.New("name is required")
	}

	if month.year == 0 {
		return errors.New("year is required")
	}

	if month.startDate.IsZero() {
		return errors.New("start date is required")
	}

	if month.endDate.IsZero() {
		return errors.New("end date is required")
	}

	if month.startDate.Equal(month.endDate) || month.startDate.After(month.endDate) {
		return errors.New("start date must be greater than the end date")
	}

	return nil
}
