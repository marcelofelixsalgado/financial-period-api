package mocks

import (
	"github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/balance/list"

	"github.com/marcelofelixsalgado/financial-commons/pkg/usecase/status"

	"github.com/stretchr/testify/mock"
)

type ListUseCaseMock struct {
	mock.Mock
}

func (m *ListUseCaseMock) Execute(input list.InputListBalanceDto) (list.OutputListBalanceDto, status.InternalStatus, error) {
	args := m.Called(input)
	return args.Get(0).(list.OutputListBalanceDto), args.Get(1).(status.InternalStatus), args.Error(2)
}
