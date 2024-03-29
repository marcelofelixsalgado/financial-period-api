package mocks

import (
	"time"

	"github.com/marcelofelixsalgado/financial-period-api/pkg/domain/period/entity"
	"github.com/marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/filter"

	"github.com/marcelofelixsalgado/financial-commons/pkg/infrastructure/repository/status"
	repositoryInternalStatus "github.com/marcelofelixsalgado/financial-commons/pkg/infrastructure/repository/status"

	"github.com/stretchr/testify/mock"
)

type PeriodRepositoryMock struct {
	mock.Mock
}

func (m *PeriodRepositoryMock) Create(period entity.IPeriod) (repositoryInternalStatus.RepositoryInternalStatus, error) {
	args := m.Called(period)
	return repositoryInternalStatus.Success, args.Error(0)
}

func (m *PeriodRepositoryMock) Update(period entity.IPeriod) (repositoryInternalStatus.RepositoryInternalStatus, error) {
	args := m.Called(period)
	return repositoryInternalStatus.Success, args.Error(0)
}

func (m *PeriodRepositoryMock) FindById(id string) (entity.IPeriod, error) {
	args := m.Called(id)
	return args.Get(0).(entity.IPeriod), args.Error(1)
}

func (m *PeriodRepositoryMock) List(filterParameters []filter.FilterParameter, tenantId string) ([]entity.IPeriod, error) {
	args := m.Called(filterParameters, tenantId)
	return args.Get(0).([]entity.IPeriod), args.Error(1)
}

func (m *PeriodRepositoryMock) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *PeriodRepositoryMock) FindOverlap(startDate time.Time, endDate time.Time, tenantId string) (status.RepositoryInternalStatus, error) {
	args := m.Called(startDate, endDate, tenantId)
	return repositoryInternalStatus.Success, args.Error(0)
}
