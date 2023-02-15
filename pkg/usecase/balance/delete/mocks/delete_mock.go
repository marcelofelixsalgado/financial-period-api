package mocks

import (
	"marcelofelixsalgado/financial-period-api/pkg/usecase/balance/delete"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/status"

	"github.com/stretchr/testify/mock"
)

type DeleteUseCaseMock struct {
	mock.Mock
}

func (m *DeleteUseCaseMock) Execute(input delete.InputDeleteBalanceDto) (delete.OutputDeleteBalanceDto, status.InternalStatus, error) {
	args := m.Called(input)
	return args.Get(0).(delete.OutputDeleteBalanceDto), args.Get(1).(status.InternalStatus), args.Error(2)
}
