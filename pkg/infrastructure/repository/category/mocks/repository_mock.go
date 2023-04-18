package mocks

import (
	"marcelofelixsalgado/financial-period-api/pkg/domain/category/entity"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/filter"

	"github.com/marcelofelixsalgado/financial-commons/pkg/infrastructure/repository/status"
	repositoryInternalStatus "github.com/marcelofelixsalgado/financial-commons/pkg/infrastructure/repository/status"

	"github.com/stretchr/testify/mock"
)

type CategoryRepositoryMock struct {
	mock.Mock
}

func (m *CategoryRepositoryMock) Create(category entity.ICategory) (status.RepositoryInternalStatus, error) {
	args := m.Called(category)
	return repositoryInternalStatus.Success, args.Error(0)
}

func (m *CategoryRepositoryMock) Update(category entity.ICategory) (status.RepositoryInternalStatus, error) {
	args := m.Called(category)
	return repositoryInternalStatus.Success, args.Error(0)
}

func (m *CategoryRepositoryMock) FindById(id string) (entity.ICategory, error) {
	args := m.Called(id)
	return args.Get(0).(entity.ICategory), args.Error(1)
}

func (m *CategoryRepositoryMock) List(filterParameters []filter.FilterParameter, tenantId string) ([]entity.ICategory, error) {
	args := m.Called(filterParameters, tenantId)
	return args.Get(0).([]entity.ICategory), args.Error(1)
}

func (m *CategoryRepositoryMock) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
