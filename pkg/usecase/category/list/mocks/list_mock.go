package mocks

import (
	"marcelofelixsalgado/financial-period-api/pkg/usecase/category/list"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/status"

	"github.com/stretchr/testify/mock"
)

type CreateUseCaseMock struct {
	mock.Mock
}

func (m *CreateUseCaseMock) Execute(input list.InputListCategoryDto) (list.OutputListCategoryDto, status.InternalStatus, error) {
	args := m.Called(input)
	return args.Get(0).(list.OutputListCategoryDto), args.Get(1).(status.InternalStatus), args.Error(2)
}
