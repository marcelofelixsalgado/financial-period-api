package find_test

import (
	"testing"
	"time"

	"github.com/marcelofelixsalgado/financial-period-api/pkg/domain/category/entity"
	transactionTypeEntity "github.com/marcelofelixsalgado/financial-period-api/pkg/domain/transactiontype/entity"
	"github.com/marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/category/mocks"
	"github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/category/find"

	useCaseStatus "github.com/marcelofelixsalgado/financial-commons/pkg/usecase/status"

	"github.com/stretchr/testify/assert"
)

func TestFindCategoryUseCase_Execute(t *testing.T) {
	m := &mocks.CategoryRepositoryMock{}

	transactionType, _ := transactionTypeEntity.NewTransactionType("EXPENSE", "Despesa")

	category, _ := entity.NewCategory("1", "11", "DF", "Despesa fixa", transactionType, time.Time{}, time.Time{})

	m.On("FindById", category.GetId()).Return(category, nil)

	useCase := find.NewFindUseCase(m)

	input := find.InputFindCategoryDto{
		Id: category.GetId(),
	}

	output, internalStatus, err := useCase.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.Id)
	assert.NotEmpty(t, output.CreatedAt)
	assert.Equal(t, category.GetName(), output.Name)
	assert.Equal(t, category.GetCode(), output.Code)
	assert.Equal(t, category.GetTransactionType().GetCode(), output.TransactionType.Code)
	assert.Equal(t, internalStatus, useCaseStatus.Success)
	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "FindById", 1)
}
