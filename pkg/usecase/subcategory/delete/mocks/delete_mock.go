package mocks

import (
	"github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/subcategory/delete"

	"github.com/marcelofelixsalgado/financial-commons/pkg/usecase/status"

	"github.com/stretchr/testify/mock"
)

type DeleteUseCaseMock struct {
	mock.Mock
}

func (m *DeleteUseCaseMock) Execute(input delete.InputDeleteSubCategoryDto) (delete.OutputDeleteSubCategoryDto, status.InternalStatus, error) {
	args := m.Called(input)
	return args.Get(0).(delete.OutputDeleteSubCategoryDto), args.Get(1).(status.InternalStatus), args.Error(2)
}
