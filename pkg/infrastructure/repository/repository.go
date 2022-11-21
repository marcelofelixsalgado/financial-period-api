package repository

import (
	"marcelofelixsalgado/financial-period-api/pkg/domain/period/entity"
	"time"
)

type IRepository interface {
	Create(entity.IPeriod) error
	Update() error
	Find(id string) (entity.Period, error)
	FindAll() ([]entity.Period, error)
	Delete(id string) error
}

type PeriodModel struct {
	id        string
	code      string
	name      string
	year      int
	startDate time.Time
	endDate   time.Time
	createdAt time.Time
	updatedAt time.Time
}
