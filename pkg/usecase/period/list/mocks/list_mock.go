package mocks

import (
	"github.com/marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/filter"
	"github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/period/list"

	"github.com/marcelofelixsalgado/financial-commons/pkg/usecase/status"

	"github.com/stretchr/testify/mock"
)

type ListUseCaseMock struct {
	mock.Mock
}

func (m *ListUseCaseMock) Execute(input list.InputListPeriodDto, filterParameters []filter.FilterParameter) (list.OutputListPeriodDto, status.InternalStatus, error) {
	args := m.Called(input, filterParameters)
	return args.Get(0).(list.OutputListPeriodDto), args.Get(1).(status.InternalStatus), args.Error(2)
}
