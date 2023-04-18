package delete_test

import (
	"marcelofelixsalgado/financial-period-api/pkg/domain/category/entity"
	transactionTypeEntity "marcelofelixsalgado/financial-period-api/pkg/domain/transactiontype/entity"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/category/mocks"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/category/delete"
	"testing"
	"time"

	useCaseStatus "github.com/marcelofelixsalgado/financial-commons/pkg/usecase/status"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDeleteCategoryUseCase_Execute(t *testing.T) {
	m := &mocks.CategoryRepositoryMock{}

	transactiontype, _ := transactionTypeEntity.NewTransactionType("EXPENSE", "Despesa")
	category, _ := entity.NewCategory("1", "11", "DF", "Despesa fixa", transactiontype, time.Time{}, time.Time{})

	m.On("FindById", category.GetId()).Return(category, nil)
	m.On("Delete", mock.Anything).Return(nil)

	useCase := delete.NewDeleteUseCase(m)

	input := delete.InputDeleteCategoryDto{
		Id: category.GetId(),
	}

	output, internalStatus, err := useCase.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, internalStatus, useCaseStatus.Success)
	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "FindById", 1)
	m.AssertNumberOfCalls(t, "Delete", 1)
}
