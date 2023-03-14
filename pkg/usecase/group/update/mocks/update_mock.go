package mocks

import (
	"marcelofelixsalgado/financial-period-api/pkg/usecase/group/update"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/status"

	"github.com/stretchr/testify/mock"
)

type UpdateUseCaseMock struct {
	mock.Mock
}

func (m *UpdateUseCaseMock) Execute(input update.InputUpdateGroupDto) (update.OutputUpdateGroupDto, status.InternalStatus, error) {
	args := m.Called(input)
	return args.Get(0).(update.OutputUpdateGroupDto), args.Get(1).(status.InternalStatus), args.Error(2)
}
