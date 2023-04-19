package mocks

import (
	"github.com/marcelofelixsalgado/financial-period-api/pkg/domain/user/entity"
	"github.com/marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/filter"

	repositoryInternalStatus "github.com/marcelofelixsalgado/financial-commons/pkg/infrastructure/repository/status"

	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (m *UserRepositoryMock) Create(user entity.IUser) (repositoryInternalStatus.RepositoryInternalStatus, error) {
	args := m.Called(user)
	return repositoryInternalStatus.Success, args.Error(0)
}

func (m *UserRepositoryMock) Update(user entity.IUser) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *UserRepositoryMock) FindById(id string) (entity.IUser, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(entity.IUser), args.Error(1)
}

func (m *UserRepositoryMock) FindByEmail(email string) (entity.IUser, error) {
	return nil, nil
}

func (m *UserRepositoryMock) List(filterParameters []filter.FilterParameter) ([]entity.IUser, error) {
	args := m.Called(filterParameters)
	return args.Get(0).([]entity.IUser), args.Error(1)
}

func (m *UserRepositoryMock) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
