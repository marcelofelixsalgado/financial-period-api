package mocks

import (
	"github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/balance/create"

	"github.com/marcelofelixsalgado/financial-commons/pkg/usecase/status"

	"github.com/stretchr/testify/mock"
)

type CreateUseCaseMock struct {
	mock.Mock
}

func (m *CreateUseCaseMock) Execute(input create.InputCreateBalanceDto) (create.OutputCreateBalanceDto, status.InternalStatus, error) {
	args := m.Called(input)
	return args.Get(0).(create.OutputCreateBalanceDto), args.Get(1).(status.InternalStatus), args.Error(2)
}
