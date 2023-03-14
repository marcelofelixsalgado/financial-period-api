package group

import (
	"marcelofelixsalgado/financial-period-api/pkg/domain/group/entity"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/filter"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/status"
)

type IGroupRepository interface {
	Create(entity.IGroup) (status.RepositoryInternalStatus, error)
	Update(entity.IGroup) (status.RepositoryInternalStatus, error)
	FindById(id string) (entity.IGroup, error)
	List(filterParameters []filter.FilterParameter, tenantId string) ([]entity.IGroup, error)
	Delete(id string) error
}
