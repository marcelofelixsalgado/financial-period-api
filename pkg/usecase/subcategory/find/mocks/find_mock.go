package mocks

import (
	"github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/subcategory/find"

	"github.com/marcelofelixsalgado/financial-commons/pkg/usecase/status"

	"github.com/stretchr/testify/mock"
)

type FindUseCaseMock struct {
	mock.Mock
}

func (m *FindUseCaseMock) Execute(input find.InputFindSubCategoryDto) (find.OutputFindSubCategoryDto, status.InternalStatus, error) {
	args := m.Called(input)
	return args.Get(0).(find.OutputFindSubCategoryDto), args.Get(1).(status.InternalStatus), args.Error(2)
}
