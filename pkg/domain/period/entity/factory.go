package entity

import (
	"time"

	"github.com/google/uuid"
)

func Create(code string, name string, year int, startDate time.Time, endDate time.Time) (IPeriod, error) {
	return NewPeriod(uuid.NewString(), code, name, year, startDate, endDate)
}
