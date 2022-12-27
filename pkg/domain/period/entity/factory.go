package entity

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

func Create(code string, name string, year int, startDate time.Time, endDate time.Time) (IPeriod, error) {
	return NewPeriod(uuid.NewV4().String(), code, name, year, startDate, endDate, time.Now(), time.Time{})
}
