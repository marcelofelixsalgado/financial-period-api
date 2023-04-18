package period

import (
	"marcelofelixsalgado/financial-period-api/pkg/domain/period/entity"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/filter"
	"time"

	"github.com/marcelofelixsalgado/financial-commons/pkg/infrastructure/repository/status"
)

type IPeriodRepository interface {
	Create(entity.IPeriod) (status.RepositoryInternalStatus, error)
	Update(entity.IPeriod) (status.RepositoryInternalStatus, error)
	FindById(id string) (entity.IPeriod, error)
	List(filterParameter []filter.FilterParameter, tenantId string) ([]entity.IPeriod, error)
	FindOverlap(startDate time.Time, endDate time.Time, tenantId string) (status.RepositoryInternalStatus, error)
	Delete(id string) error
}
