package mocks

import (
	"marcelofelixsalgado/financial-period-api/pkg/usecase/status"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/subcategory/delete"

	"github.com/stretchr/testify/mock"
)

type DeleteUseCaseMock struct {
	mock.Mock
}

func (m *DeleteUseCaseMock) Execute(input delete.InputDeleteUseCaseDto) (delete.OutputDeleteUseCaseDto, status.InternalStatus, error) {
	args := m.Called(input)
	return args.Get(0).(delete.OutputDeleteUseCaseDto), args.Get(1).(status.InternalStatus), args.Error(2)
}
