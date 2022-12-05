package user

import (
	"marcelofelixsalgado/financial-period-api/pkg/domain/user/entity"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/filter"
)

type IRepository interface {
	Create(entity.IUser) error
	Update(entity.IUser) error
	Find(id string) (entity.IUser, error)
	List([]filter.FilterParameter) ([]entity.IUser, error)
	Delete(id string) error
}
