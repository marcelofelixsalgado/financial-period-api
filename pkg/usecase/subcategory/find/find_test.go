package find_test

import (
	"testing"
	"time"

	categoryEntity "github.com/marcelofelixsalgado/financial-period-api/pkg/domain/category/entity"
	"github.com/marcelofelixsalgado/financial-period-api/pkg/domain/subcategory/entity"
	transactionTypeEntity "github.com/marcelofelixsalgado/financial-period-api/pkg/domain/transactiontype/entity"
	"github.com/marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/subcategory/mocks"
	"github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/subcategory/find"

	useCaseStatus "github.com/marcelofelixsalgado/financial-commons/pkg/usecase/status"

	"github.com/stretchr/testify/assert"
)

func TestFindSubCategoryUseCase_Execute(t *testing.T) {
	m := &mocks.SubCategoryRepositoryMock{}

	tenantId := "11"
	transactionType, _ := transactionTypeEntity.NewTransactionType("EXPENSE", "Despesa")
	category, _ := categoryEntity.NewCategory("1", tenantId, "DF", "Despesa fixa", transactionType, time.Time{}, time.Time{})
	subCategory, _ := entity.NewSubCategory("111", tenantId, "CO", "Condimínio", category, time.Time{}, time.Time{})

	m.On("FindById", subCategory.GetId()).Return(subCategory, nil)

	useCase := find.NewFindUseCase(m)

	input := find.InputFindSubCategoryDto{
		Id: subCategory.GetId(),
	}

	output, internalStatus, err := useCase.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.Id)
	assert.NotEmpty(t, output.CreatedAt)
	assert.Equal(t, subCategory.GetName(), output.Name)
	assert.Equal(t, subCategory.GetCode(), output.Code)
	assert.Equal(t, subCategory.GetCategory().GetId(), output.Category.Id)
	assert.Equal(t, subCategory.GetCategory().GetCode(), output.Category.Code)
	assert.Equal(t, subCategory.GetCategory().GetName(), output.Category.Name)
	assert.Equal(t, subCategory.GetCategory().GetTransactionType().GetCode(), output.Category.TransactionType.Code)
	assert.Equal(t, subCategory.GetCategory().GetTransactionType().GetName(), output.Category.TransactionType.Name)
	assert.Equal(t, internalStatus, useCaseStatus.Success)
	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "FindById", 1)
}
