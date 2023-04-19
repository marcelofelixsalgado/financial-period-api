package mocks

import (
	"github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/subcategory/update"

	"github.com/marcelofelixsalgado/financial-commons/pkg/usecase/status"

	"github.com/stretchr/testify/mock"
)

type UpdateUseCaseMock struct {
	mock.Mock
}

func (m *UpdateUseCaseMock) Execute(input update.InputUpdateSubCategoryDto) (update.OutputUpdateSubCategoryDto, status.InternalStatus, error) {
	args := m.Called(input)
	return args.Get(0).(update.OutputUpdateSubCategoryDto), args.Get(1).(status.InternalStatus), args.Error(2)
}
