package mocks

import (
	"marcelofelixsalgado/financial-period-api/pkg/usecase/transactiontype/find"

	"github.com/marcelofelixsalgado/financial-commons/pkg/usecase/status"

	"github.com/stretchr/testify/mock"
)

type FindUseCaseMock struct {
	mock.Mock
}

func (m *FindUseCaseMock) Execute(input find.InputFindTransactionTypeDto) (find.OutputFindTransactionTypeDto, status.InternalStatus, error) {
	args := m.Called(input)
	return args.Get(0).(find.OutputFindTransactionTypeDto), args.Get(1).(status.InternalStatus), args.Error(2)
}
