package mocks

import (
	"github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/category/delete"

	"github.com/marcelofelixsalgado/financial-commons/pkg/usecase/status"

	"github.com/stretchr/testify/mock"
)

type DeleteUseCaseMock struct {
	mock.Mock
}

func (m *DeleteUseCaseMock) Execute(input delete.InputDeleteCategoryDto) (delete.OutputDeleteCategoryDto, status.InternalStatus, error) {
	args := m.Called(input)
	return args.Get(0).(delete.OutputDeleteCategoryDto), args.Get(1).(status.InternalStatus), args.Error(2)
}
