package category

import (
	"marcelofelixsalgado/financial-period-api/pkg/domain/category/entity"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/filter"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/status"
)

type ICategoryRepository interface {
	Create(entity.ICategory) (status.RepositoryInternalStatus, error)
	Update(entity.ICategory) (status.RepositoryInternalStatus, error)
	FindById(id string) (entity.ICategory, error)
	List(filterParameters []filter.FilterParameter, tenantId string) ([]entity.ICategory, error)
	Delete(id string) error
}
