package subcategory

import (
	"marcelofelixsalgado/financial-period-api/pkg/domain/subcategory/entity"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/filter"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/status"
)

type ISubCategoryRepository interface {
	Create(entity entity.ISubCategory) (status.RepositoryInternalStatus, error)
	Update(entity entity.ISubCategory) (status.RepositoryInternalStatus, error)
	FindById(id string) (entity.ISubCategory, error)
	List(filterParameters []filter.FilterParameter, tenantId string) ([]entity.ISubCategory, error)
	Delete(id string) error
}
