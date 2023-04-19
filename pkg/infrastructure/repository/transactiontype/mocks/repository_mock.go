package mocks

import (
	"github.com/marcelofelixsalgado/financial-period-api/pkg/domain/transactiontype/entity"
	"github.com/marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/filter"

	"github.com/stretchr/testify/mock"
)

type TransactionTypeRepositoryMock struct {
	mock.Mock
}

func (m *TransactionTypeRepositoryMock) FindByCode(code string) (entity.ITransactionType, error) {
	args := m.Called(code)
	return args.Get(0).(entity.ITransactionType), args.Error(1)
}

func (m *TransactionTypeRepositoryMock) List(filterParameters []filter.FilterParameter) ([]entity.ITransactionType, error) {
	args := m.Called(filterParameters)
	return args.Get(0).([]entity.ITransactionType), args.Error(1)
}
