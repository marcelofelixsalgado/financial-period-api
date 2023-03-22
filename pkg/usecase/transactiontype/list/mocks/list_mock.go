package mocks

import (
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/filter"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/status"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/transactiontype/list"

	"github.com/stretchr/testify/mock"
)

type ListUseCaseMock struct {
	mock.Mock
}

func (m *ListUseCaseMock) Execute(input list.InputListTransactionTypeDto, filterParameters []filter.FilterParameter) (list.OutputListTransactionTypeDto, status.InternalStatus, error) {
	args := m.Called(input, filterParameters)
	return args.Get(0).(list.OutputListTransactionTypeDto), args.Get(1).(status.InternalStatus), args.Error(2)
}
