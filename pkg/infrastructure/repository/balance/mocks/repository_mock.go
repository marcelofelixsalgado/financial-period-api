package mocks

import (
	"marcelofelixsalgado/financial-period-api/pkg/domain/balance/entity"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/status"

	repositoryInternalStatus "marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/status"

	"github.com/stretchr/testify/mock"
)

type BalanceRepositoryMock struct {
	mock.Mock
}

func (m *BalanceRepositoryMock) Create(balance entity.IBalance) (status.RepositoryInternalStatus, error) {
	args := m.Called(balance)
	return repositoryInternalStatus.Success, args.Error(0)
}

func (m *BalanceRepositoryMock) Update(balance entity.IBalance) error {
	args := m.Called(balance)
	return args.Error(0)
}

func (m *BalanceRepositoryMock) FindById(id string) (entity.IBalance, error) {
	args := m.Called(id)
	return args.Get(0).(entity.IBalance), args.Error(1)
}

func (m *BalanceRepositoryMock) List(tenantId string, periodId string) ([]entity.IBalance, error) {
	args := m.Called(tenantId, periodId)
	return args.Get(0).([]entity.IBalance), args.Error(1)
}

func (m *BalanceRepositoryMock) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
