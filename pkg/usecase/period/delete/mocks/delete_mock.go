package mocks

import (
	"marcelofelixsalgado/financial-period-api/pkg/usecase/period/delete"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/status"

	"github.com/stretchr/testify/mock"
)

type DeleteUseCaseMock struct {
	mock.Mock
}

func (m *DeleteUseCaseMock) Execute(input delete.InputDeletePeriodDto) (delete.OutputDeletePeriodDto, status.InternalStatus, error) {
	args := m.Called(input)
	return args.Get(0).(delete.OutputDeletePeriodDto), args.Get(1).(status.InternalStatus), args.Error(2)
}
