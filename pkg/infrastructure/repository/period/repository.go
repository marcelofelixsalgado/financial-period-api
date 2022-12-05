package period

import (
	"marcelofelixsalgado/financial-period-api/pkg/domain/period/entity"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/filter"
)

type IRepository interface {
	Create(entity.IPeriod) error
	Update(entity.IPeriod) error
	FindById(id string) (entity.IPeriod, error)
	List([]filter.FilterParameter) ([]entity.IPeriod, error)
	Delete(id string) error
}
