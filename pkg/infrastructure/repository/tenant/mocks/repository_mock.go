package mocks

import (
	"marcelofelixsalgado/financial-period-api/pkg/domain/tenant/entity"

	repositoryInternalStatus "github.com/marcelofelixsalgado/financial-commons/pkg/infrastructure/repository/status"

	"github.com/stretchr/testify/mock"
)

type TenantRepositoryMock struct {
	mock.Mock
}

func (m *TenantRepositoryMock) Create(tenant entity.ITenant) (repositoryInternalStatus.RepositoryInternalStatus, error) {
	args := m.Called(tenant)
	return repositoryInternalStatus.Success, args.Error(0)
}
