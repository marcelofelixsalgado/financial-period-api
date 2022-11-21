package entity

import (
	"errors"
	"time"
)

type IPeriod interface {
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

type Period struct {
	id        string
	code      string
	name      string
	year      int
	startDate time.Time
	endDate   time.Time
	createdAt time.Time
	updatedAt time.Time
}

func (period Period) GetId() string {
	return period.id
}

func (period Period) GetCode() string {
	return period.code
}

func (period Period) GetName() string {
	return period.name
}

func (period Period) GetYear() int {
	return period.year
}

func (period Period) GetStartDate() time.Time {
	return period.startDate
}

func (period Period) GetEndDate() time.Time {
	return period.endDate
}

func (period Period) GetCreatedAt() time.Time {
	return period.createdAt
}

func (period Period) GetUpdatedAt() time.Time {
	return period.updatedAt
}

func (period Period) SetCode(code string) {
	period.code = code
	validate(period)
}

func (period Period) SetName(name string) {
	period.name = name
	validate(period)
}

func (period Period) SetYear(year int) {
	period.year = year
	validate(period)
}

func (period Period) SetStartDate(startDate time.Time) {
	period.startDate = startDate
	validate(period)
}

func (period Period) SetEndDate(endDate time.Time) {
	period.endDate = endDate
	validate(period)
}

func (period Period) SetUpdatedAt(updatedAt time.Time) {
	period.updatedAt = updatedAt
	validate(period)
}

func NewPeriod(id string, code string, name string, year int, startDate time.Time, endDate time.Time) (IPeriod, error) {
	period := Period{
		id:        id,
		code:      code,
		name:      name,
		year:      year,
		startDate: startDate,
		endDate:   endDate,
		createdAt: time.Now(),
	}
	if err := validate(period); err != nil {
		return nil, err
	}
	return period, nil
}

func validate(period Period) error {
	if period.id == "" {
		return errors.New("id is required")
	}

	if period.code == "" {
		return errors.New("code is required")
	}

	if period.name == "" {
		return errors.New("name is required")
	}

	if period.year == 0 {
		return errors.New("year is required")
	}

	if period.startDate.IsZero() {
		return errors.New("start date is required")
	}

	if period.endDate.IsZero() {
		return errors.New("end date is required")
	}

	if period.startDate.Equal(period.endDate) || period.startDate.After(period.endDate) {
		return errors.New("start date must be greater than the end date")
	}

	return nil
}
