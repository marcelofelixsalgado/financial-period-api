package mocks

import (
	"github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/period/findbyid"

	"github.com/marcelofelixsalgado/financial-commons/pkg/usecase/status"

	"github.com/stretchr/testify/mock"
)

type FindByIdUseCaseMock struct {
	mock.Mock
}

func (m *FindByIdUseCaseMock) Execute(input findbyid.InputFindByIdPeriodDto) (findbyid.OutputFindByIdPeriodDto, status.InternalStatus, error) {
	args := m.Called(input)
	return args.Get(0).(findbyid.OutputFindByIdPeriodDto), args.Get(1).(status.InternalStatus), args.Error(2)
}
