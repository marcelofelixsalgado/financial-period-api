package create_test

import (
	"testing"
	"time"

	"github.com/marcelofelixsalgado/financial-period-api/pkg/domain/category/entity"
	transactionTypeEntity "github.com/marcelofelixsalgado/financial-period-api/pkg/domain/transactiontype/entity"
	categoryRepositoryMock "github.com/marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/category/mocks"
	subCategoryRepositoryMock "github.com/marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/subcategory/mocks"
	"github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/subcategory/create"

	useCaseStatus "github.com/marcelofelixsalgado/financial-commons/pkg/usecase/status"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateSubCategoryUseCase_Execute(t *testing.T) {

	categoryRepositoryMock := &categoryRepositoryMock.CategoryRepositoryMock{}
	subcategoryRepositoryMock := &subCategoryRepositoryMock.SubCategoryRepositoryMock{}

	tenantId := "11"
	transactionType, _ := transactionTypeEntity.NewTransactionType("EXPENSE", "Expenses")
	category, _ := entity.NewCategory("1", tenantId, "DF", "Despesa fixa", transactionType, time.Time{}, time.Time{})

	categoryRepositoryMock.On("FindById", category.GetId()).Return(category, nil)
	subcategoryRepositoryMock.On("Create", mock.Anything).Return(nil)

	useCase := create.NewCreateUseCase(subcategoryRepositoryMock, categoryRepositoryMock)

	input := create.InputCreateSubCategoryDto{
		TenantId: tenantId,
		Code:     "TR",
		Name:     "Transporte",
		Category: create.CategoryInput{
			Id: category.GetId(),
		},
	}

	output, internalStatus, err := useCase.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.Id)
	assert.NotEmpty(t, output.CreatedAt)
	assert.Equal(t, input.Code, output.Code)
	assert.Equal(t, input.Name, output.Name)
	assert.Equal(t, category.GetId(), output.Category.Id)
	assert.Equal(t, category.GetCode(), output.Category.Code)
	assert.Equal(t, category.GetTransactionType().GetCode(), output.Category.TransactionType.Code)
	assert.Equal(t, category.GetTransactionType().GetName(), output.Category.TransactionType.Name)
	assert.Equal(t, internalStatus, useCaseStatus.Success)
	subcategoryRepositoryMock.AssertExpectations(t)
	subcategoryRepositoryMock.AssertNumberOfCalls(t, "Create", 1)
}
