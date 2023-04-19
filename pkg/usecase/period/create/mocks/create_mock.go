package mocks

import (
	"github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/period/create"

	"github.com/marcelofelixsalgado/financial-commons/pkg/usecase/status"
	"github.com/stretchr/testify/mock"
)

type CreateUseCaseMock struct {
	mock.Mock
}

func (m *CreateUseCaseMock) Execute(input create.InputCreatePeriodDto) (create.OutputCreatePeriodDto, status.InternalStatus, error) {
	args := m.Called(input)
	return args.Get(0).(create.OutputCreatePeriodDto), args.Get(1).(status.InternalStatus), args.Error(2)
}
