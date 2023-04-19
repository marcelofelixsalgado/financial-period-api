package user

import (
	"github.com/marcelofelixsalgado/financial-period-api/pkg/domain/user/entity"
	"github.com/marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/filter"

	"github.com/marcelofelixsalgado/financial-commons/pkg/infrastructure/repository/status"
)

type IUserRepository interface {
	Create(entity.IUser) (status.RepositoryInternalStatus, error)
	Update(entity.IUser) error
	FindById(id string) (entity.IUser, error)
	FindByEmail(email string) (entity.IUser, error)
	List(filterParameters []filter.FilterParameter) ([]entity.IUser, error)
	Delete(id string) error
}
