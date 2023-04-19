package mocks

import (
	"github.com/marcelofelixsalgado/financial-period-api/pkg/domain/subcategory/entity"
	"github.com/marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/filter"

	"github.com/marcelofelixsalgado/financial-commons/pkg/infrastructure/repository/status"
	repositoryInternalStatus "github.com/marcelofelixsalgado/financial-commons/pkg/infrastructure/repository/status"

	"github.com/stretchr/testify/mock"
)

type SubCategoryRepositoryMock struct {
	mock.Mock
}

func (m *SubCategoryRepositoryMock) Create(entity entity.ISubCategory) (status.RepositoryInternalStatus, error) {
	args := m.Called(entity)
	return repositoryInternalStatus.Success, args.Error(0)
}

func (m *SubCategoryRepositoryMock) Update(entity entity.ISubCategory) (status.RepositoryInternalStatus, error) {
	args := m.Called(entity)
	return repositoryInternalStatus.Success, args.Error(0)
}

func (m *SubCategoryRepositoryMock) FindById(id string) (entity.ISubCategory, error) {
	args := m.Called(id)
	return args.Get(0).(entity.ISubCategory), args.Error(1)
}

func (m *SubCategoryRepositoryMock) List(filterParameters []filter.FilterParameter, tenantId string) ([]entity.ISubCategory, error) {
	args := m.Called(filterParameters, tenantId)
	return args.Get(0).([]entity.ISubCategory), args.Error(1)
}

func (m *SubCategoryRepositoryMock) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
