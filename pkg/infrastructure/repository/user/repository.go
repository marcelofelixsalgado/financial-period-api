package user

import (
	"marcelofelixsalgado/financial-period-api/pkg/domain/user/entity"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/filter"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/status"
)

type IUserRepository interface {
	Create(entity.IUser) (status.RepositoryInternalStatus, error)
	Update(entity.IUser) error
	FindById(id string) (entity.IUser, error)
	FindByEmail(email string) (entity.IUser, error)
	List([]filter.FilterParameter) ([]entity.IUser, error)
	Delete(id string) error
}
