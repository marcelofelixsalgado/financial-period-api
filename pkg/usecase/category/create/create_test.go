package create_test

import (
	"marcelofelixsalgado/financial-period-api/pkg/domain/transactiontype/entity"
	categoryRepositoryMock "marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/category/mocks"
	transactionTypeRepositoryMock "marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/transactiontype/mocks"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/category/create"
	"testing"

	useCaseStatus "github.com/marcelofelixsalgado/financial-commons/pkg/usecase/status"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateCategoryUseCase_Execute(t *testing.T) {
	categoryRepositoryMock := &categoryRepositoryMock.CategoryRepositoryMock{}
	transactionTypeRepositoryMock := &transactionTypeRepositoryMock.TransactionTypeRepositoryMock{}

	transactionType, _ := entity.NewTransactionType("EXPENSE", "Expenses")

	transactionTypeRepositoryMock.On("FindByCode", transactionType.GetCode()).Return(transactionType, nil)
	categoryRepositoryMock.On("Create", mock.Anything).Return(nil)

	useCase := create.NewCreateUseCase(categoryRepositoryMock, transactionTypeRepositoryMock)

	input := create.InputCreateCategoryDto{
		TenantId: "123",
		Code:     "DF",
		Name:     "Despesa fixa",
		TransactionType: create.TransactionTypeInput{
			Code: transactionType.GetCode(),
		},
	}

	output, internalStatus, err := useCase.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.Id)
	assert.NotEmpty(t, output.CreatedAt)
	assert.Equal(t, input.Code, output.Code)
	assert.Equal(t, input.Name, output.Name)
	assert.Equal(t, input.TransactionType.Code, output.TransactionType.Code)
	assert.Equal(t, internalStatus, useCaseStatus.Success)
	categoryRepositoryMock.AssertExpectations(t)
	categoryRepositoryMock.AssertNumberOfCalls(t, "Create", 1)

}
