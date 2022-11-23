package repository

import (
	"marcelofelixsalgado/financial-period-api/pkg/domain/period/entity"
)

type IRepository interface {
	Create(entity.IPeriod) error
	Update(entity.IPeriod) error
	FindById(id string) (entity.IPeriod, error)
	FindAll() ([]entity.IPeriod, error)
	Delete(id string) error
}
