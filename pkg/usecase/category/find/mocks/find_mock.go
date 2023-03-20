package mocks

import (
	"marcelofelixsalgado/financial-period-api/pkg/usecase/category/find"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/status"

	"github.com/stretchr/testify/mock"
)

type FindUseCaseMock struct {
	mock.Mock
}

func (m *FindUseCaseMock) Execute(input find.InputFindCategoryDto) (find.OutputFindCategoryDto, status.InternalStatus, error) {
	args := m.Called(input)
	return args.Get(0).(find.OutputFindCategoryDto), args.Get(1).(status.InternalStatus), args.Error(2)
}
