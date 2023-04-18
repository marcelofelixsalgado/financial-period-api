package mocks

import (
	"marcelofelixsalgado/financial-period-api/pkg/usecase/credentials/create"

	"github.com/marcelofelixsalgado/financial-commons/pkg/usecase/status"

	"github.com/stretchr/testify/mock"
)

type CreateUseCaseMock struct {
	mock.Mock
}

func (m *CreateUseCaseMock) Execute(input create.InputCreateUserCredentialsDto) (create.OutputCreateUserCredentialsDto, status.InternalStatus, error) {
	args := m.Called(input)
	return args.Get(0).(create.OutputCreateUserCredentialsDto), args.Get(1).(status.InternalStatus), args.Error(2)
}
