package list_test

import (
	categoryEntity "marcelofelixsalgado/financial-period-api/pkg/domain/category/entity"
	"marcelofelixsalgado/financial-period-api/pkg/domain/subcategory/entity"
	transactionTypeEntity "marcelofelixsalgado/financial-period-api/pkg/domain/transactiontype/entity"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/filter"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/subcategory/mocks"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/subcategory/list"
	"testing"
	"time"

	useCaseStatus "marcelofelixsalgado/financial-period-api/pkg/usecase/status"

	"github.com/stretchr/testify/assert"
)

func TestListSubCategoriesUseCase_Execute(t *testing.T) {

	m := &mocks.SubCategoryRepositoryMock{}

	tenantId := "11"

	transactionType, _ := transactionTypeEntity.NewTransactionType("EXPENSE", "Despesa")
	category1, _ := categoryEntity.NewCategory("1", tenantId, "DF", "Despesa fixa", transactionType, time.Time{}, time.Time{})
	category2, _ := categoryEntity.NewCategory("2", tenantId, "DV", "Despesa variável", transactionType, time.Time{}, time.Time{})

	subCategory1, _ := entity.NewSubCategory("111", tenantId, "CO", "Condimínio", category1, time.Time{}, time.Time{})
	subCategory2, _ := entity.NewSubCategory("222", tenantId, "TR", "Transportes", category2, time.Time{}, time.Time{})
	subCategories := []entity.ISubCategory{subCategory1, subCategory2}

	m.On("List", []filter.FilterParameter{}, tenantId).Return(subCategories, nil)

	useCase := list.NewListUseCase(m)

	input := list.InputListSubCategoryDto{
		TenantId: tenantId,
	}

	output, internalStatus, err := useCase.Execute(input, []filter.FilterParameter{})

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Len(t, output.SubCategories, 2)

	assert.NotEmpty(t, output.SubCategories[0].Id)
	assert.Equal(t, subCategories[0].GetName(), output.SubCategories[0].Name)
	assert.Equal(t, subCategories[0].GetCode(), output.SubCategories[0].Code)
	assert.Equal(t, subCategories[0].GetCategory().GetId(), output.SubCategories[0].Category.Id)
	assert.Equal(t, subCategories[0].GetCategory().GetCode(), output.SubCategories[0].Category.Code)
	assert.Equal(t, subCategories[0].GetCategory().GetName(), output.SubCategories[0].Category.Name)
	assert.Equal(t, subCategories[0].GetCategory().GetTransactionType().GetCode(), output.SubCategories[0].Category.TransactionType.Code)
	assert.Equal(t, subCategories[0].GetCategory().GetTransactionType().GetName(), output.SubCategories[0].Category.TransactionType.Name)

	assert.Equal(t, subCategories[1].GetName(), output.SubCategories[1].Name)
	assert.Equal(t, subCategories[1].GetCode(), output.SubCategories[1].Code)
	assert.Equal(t, subCategories[1].GetCategory().GetId(), output.SubCategories[1].Category.Id)
	assert.Equal(t, subCategories[1].GetCategory().GetCode(), output.SubCategories[1].Category.Code)
	assert.Equal(t, subCategories[1].GetCategory().GetName(), output.SubCategories[1].Category.Name)
	assert.Equal(t, subCategories[1].GetCategory().GetTransactionType().GetCode(), output.SubCategories[1].Category.TransactionType.Code)
	assert.Equal(t, subCategories[1].GetCategory().GetTransactionType().GetName(), output.SubCategories[1].Category.TransactionType.Name)

	assert.Equal(t, internalStatus, useCaseStatus.Success)
	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "List", 1)
}
