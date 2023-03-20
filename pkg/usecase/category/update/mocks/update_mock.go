package mocks

import (
	"marcelofelixsalgado/financial-period-api/pkg/usecase/category/update"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/status"

	"github.com/stretchr/testify/mock"
)

type CreateUseCaseMock struct {
	mock.Mock
}

func (m *CreateUseCaseMock) Execute(input update.InputUpdateCategoryDto) (update.OutputUpdateCategoryDto, status.InternalStatus, error) {
	args := m.Called(input)
	return args.Get(0).(update.OutputUpdateCategoryDto), args.Get(1).(status.InternalStatus), args.Error(2)
}
