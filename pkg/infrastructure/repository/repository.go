package repository

import (
	"marcelofelixsalgado/financial-month-api/pkg/domain/month/entity"
	"time"
)

type IRepository interface {
	Create(entity.IMonth) error
	Update() error
	Find(id string) (entity.Month, error)
	FindAll() ([]entity.Month, error)
	Delete(id string) error
}

type MonthModel struct {
	id        string
	code      string
	name      string
	year      int
	startDate time.Time
	endDate   time.Time
	createdAt time.Time
	updatedAt time.Time
}
