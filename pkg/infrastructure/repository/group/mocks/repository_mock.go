package mocks

import (
	"marcelofelixsalgado/financial-period-api/pkg/domain/group/entity"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/filter"

	repositoryInternalStatus "marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/status"

	"github.com/stretchr/testify/mock"
)

type GroupRepositoryMock struct {
	mock.Mock
}

func (m *GroupRepositoryMock) Create(group entity.IGroup) (repositoryInternalStatus.RepositoryInternalStatus, error) {
	args := m.Called(group)
	return repositoryInternalStatus.Success, args.Error(0)
}

func (m *GroupRepositoryMock) Update(group entity.IGroup) (repositoryInternalStatus.RepositoryInternalStatus, error) {
	args := m.Called(group)
	return repositoryInternalStatus.Success, args.Error(0)
}

func (m *GroupRepositoryMock) FindById(id string) (entity.IGroup, error) {
	args := m.Called(id)
	return args.Get(0).(entity.IGroup), args.Error(1)
}

func (m *GroupRepositoryMock) List(filterParameters []filter.FilterParameter, tenantId string) ([]entity.IGroup, error) {
	args := m.Called(filterParameters, tenantId)
	return args.Get(0).([]entity.IGroup), args.Error(1)
}

func (m *GroupRepositoryMock) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
