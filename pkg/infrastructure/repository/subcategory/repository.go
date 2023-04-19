package subcategory

import (
	"github.com/marcelofelixsalgado/financial-period-api/pkg/domain/subcategory/entity"
	"github.com/marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/filter"

	"github.com/marcelofelixsalgado/financial-commons/pkg/infrastructure/repository/status"
)

type ISubCategoryRepository interface {
	Create(entity entity.ISubCategory) (status.RepositoryInternalStatus, error)
	Update(entity entity.ISubCategory) (status.RepositoryInternalStatus, error)
	FindById(id string) (entity.ISubCategory, error)
	List(filterParameters []filter.FilterParameter, tenantId string) ([]entity.ISubCategory, error)
	Delete(id string) error
}
