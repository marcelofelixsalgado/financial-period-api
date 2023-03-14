package mocks

import (
	"marcelofelixsalgado/financial-period-api/pkg/usecase/group/create"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/status"

	"github.com/stretchr/testify/mock"
)

type CreateUseCaseMock struct {
	mock.Mock
}

func (m *CreateUseCaseMock) Execute(input create.InputCreateGroupDto) (create.OutputCreateGroupDto, status.InternalStatus, error) {
	args := m.Called(input)
	return args.Get(0).(create.OutputCreateGroupDto), args.Get(1).(status.InternalStatus), args.Error(2)
}
