package list_test

import (
	"marcelofelixsalgado/financial-period-api/pkg/domain/category/entity"
	transactionType "marcelofelixsalgado/financial-period-api/pkg/domain/transactiontype/entity"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/category/mocks"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/filter"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/category/list"
	useCaseStatus "marcelofelixsalgado/financial-period-api/pkg/usecase/status"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestListCategoryUseCase_Execute(t *testing.T) {
	m := &mocks.CategoryRepositoryMock{}

	tenantId := "11"

	transactionType1, _ := transactionType.NewTransactionType("EXPENSE", "Despesa")
	transactionType2, _ := transactionType.NewTransactionType("EARNING", "Receita")

	category1, _ := entity.NewCategory("1", tenantId, "DF", "Despesa fixa", transactionType1, time.Time{}, time.Time{})
	category2, _ := entity.NewCategory("2", tenantId, "SL", "Salario", transactionType2, time.Time{}, time.Time{})

	categories := []entity.ICategory{category1, category2}

	m.On("List", []filter.FilterParameter{}, tenantId).Return(categories, nil)

	useCase := list.NewListUseCase(m)

	input := list.InputListCategoryDto{
		TenantId: tenantId,
	}

	output, internalStatus, err := useCase.Execute(input, []filter.FilterParameter{})

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Len(t, output.Categories, 2)

	assert.NotEmpty(t, output.Categories[0].Id)
	assert.Equal(t, categories[0].GetName(), output.Categories[0].Name)
	assert.Equal(t, categories[0].GetCode(), output.Categories[0].Code)
	assert.Equal(t, categories[0].GetTransactionType().GetCode(), output.Categories[0].TransactionType.Code)

	assert.Equal(t, categories[1].GetName(), output.Categories[1].Name)
	assert.Equal(t, categories[1].GetCode(), output.Categories[1].Code)
	assert.Equal(t, categories[1].GetTransactionType().GetCode(), output.Categories[1].TransactionType.Code)

	assert.Equal(t, internalStatus, useCaseStatus.Success)
	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "List", 1)

}
