package mocks

import (
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/filter"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/subcategory/list"

	"github.com/marcelofelixsalgado/financial-commons/pkg/usecase/status"

	"github.com/stretchr/testify/mock"
)

type ListUseCaseMock struct {
	mock.Mock
}

func (m *ListUseCaseMock) Execute(input list.InputListSubCategoryDto, filterParameters []filter.FilterParameter) (list.OutputListSubCategoryDto, status.InternalStatus, error) {
	args := m.Called(input, filterParameters)
	return args.Get(0).(list.OutputListSubCategoryDto), args.Get(1).(status.InternalStatus), args.Error(2)
}
