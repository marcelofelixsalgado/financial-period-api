package mocks

import (
	"marcelofelixsalgado/financial-period-api/pkg/usecase/group/find"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/status"

	"github.com/stretchr/testify/mock"
)

type FindUseCaseMock struct {
	mock.Mock
}

func (m *FindUseCaseMock) Execute(input find.InputFindGroupDto) (find.OutputFindGroupDto, status.InternalStatus, error) {
	args := m.Called(input)
	return args.Get(0).(find.OutputFindGroupDto), args.Get(1).(status.InternalStatus), args.Error(2)
}
