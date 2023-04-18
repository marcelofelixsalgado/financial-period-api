package update_test

import (
	categoryEntity "marcelofelixsalgado/financial-period-api/pkg/domain/category/entity"
	subCategoryEntity "marcelofelixsalgado/financial-period-api/pkg/domain/subcategory/entity"
	transactionTypeEntity "marcelofelixsalgado/financial-period-api/pkg/domain/transactiontype/entity"
	categoryRepository "marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/category/mocks"
	subCategoryRepository "marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/subcategory/mocks"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/subcategory/update"
	"testing"
	"time"

	useCaseStatus "github.com/marcelofelixsalgado/financial-commons/pkg/usecase/status"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdateSubCategoryUseCase_Execute(t *testing.T) {

	categoryRepository := &categoryRepository.CategoryRepositoryMock{}
	subCategoryRepository := &subCategoryRepository.SubCategoryRepositoryMock{}

	tenantId := "11"
	transactiontype, _ := transactionTypeEntity.NewTransactionType("EXPENSE", "Despesa")
	category, _ := categoryEntity.NewCategory("1", tenantId, "DF", "Despesa Fixa", transactiontype, time.Time{}, time.Time{})
	subcategory, _ := subCategoryEntity.NewSubCategory("1", tenantId, "TR", "Transporte", category, time.Time{}, time.Time{})

	categoryRepository.On("FindById", category.GetId()).Return(category, nil)
	subCategoryRepository.On("FindById", subcategory.GetId()).Return(subcategory, nil)
	subCategoryRepository.On("Update", mock.Anything).Return(nil)

	useCase := update.NewUpdateUseCase(subCategoryRepository, categoryRepository)

	input := update.InputUpdateSubCategoryDto{
		Id:       subcategory.GetId(),
		TenantId: tenantId,
		Code:     subcategory.GetCode(),
		Name:     subcategory.GetName(),
		Category: update.CategoryInput{
			Id: category.GetId(),
		},
	}

	output, internalStatus, err := useCase.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.Id)
	assert.NotEmpty(t, output.CreatedAt)
	assert.Equal(t, subcategory.GetName(), output.Name)
	assert.Equal(t, subcategory.GetCode(), output.Code)
	assert.Equal(t, subcategory.GetCategory().GetCode(), output.Category.Code)
	assert.Equal(t, subcategory.GetCategory().GetName(), output.Category.Name)
	assert.Equal(t, subcategory.GetCategory().GetTransactionType().GetCode(), output.Category.TransactionType.Code)
	assert.Equal(t, subcategory.GetCategory().GetTransactionType().GetName(), output.Category.TransactionType.Name)
	assert.Equal(t, internalStatus, useCaseStatus.Success)
	subCategoryRepository.AssertExpectations(t)
	categoryRepository.AssertNumberOfCalls(t, "FindById", 1)
	subCategoryRepository.AssertNumberOfCalls(t, "Update", 1)
}
