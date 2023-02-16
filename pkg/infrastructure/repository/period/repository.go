package period

import (
	"marcelofelixsalgado/financial-period-api/pkg/domain/period/entity"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/filter"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/status"
)

type IPeriodRepository interface {
	Create(entity.IPeriod) (status.RepositoryInternalStatus, error)
	Update(entity.IPeriod) (status.RepositoryInternalStatus, error)
	FindById(id string) (entity.IPeriod, error)
	List(filterParameter []filter.FilterParameter, tenantId string) ([]entity.IPeriod, error)
	Delete(id string) error
}
