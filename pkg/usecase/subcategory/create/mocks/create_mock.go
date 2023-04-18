package mocks

import (
	"marcelofelixsalgado/financial-period-api/pkg/usecase/subcategory/create"

	"github.com/marcelofelixsalgado/financial-commons/pkg/usecase/status"

	"github.com/stretchr/testify/mock"
)

type CreateUseCaseMock struct {
	mock.Mock
}

func (m *CreateUseCaseMock) Execute(input create.InputCreateSubCategoryDto) (create.OutputCreateSubCategoryDto, status.InternalStatus, error) {
	args := m.Called(input)
	return args.Get(0).(create.OutputCreateSubCategoryDto), args.Get(1).(status.InternalStatus), args.Error(2)
}
