package update_test

import (
	"marcelofelixsalgado/financial-period-api/pkg/domain/category/entity"
	transactionType "marcelofelixsalgado/financial-period-api/pkg/domain/transactiontype/entity"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/category/mocks"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/category/update"
	useCaseStatus "marcelofelixsalgado/financial-period-api/pkg/usecase/status"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdateCategoryUseCase_Execute(t *testing.T) {
	m := &mocks.CategoryRepositoryMock{}

	transactionType, _ := transactionType.NewTransactionType("EXPENSE", "Despesas")

	category, _ := entity.NewCategory("1", "11", "DF", "Despesa fixa", transactionType, time.Now(), time.Time{})

	m.On("FindById", category.GetId()).Return(category, nil)
	m.On("Update", mock.Anything).Return(nil)

	useCase := update.NewUpdateUseCase(m)

	input := update.InputUpdateCategoryDto{
		Id:       category.GetId(),
		TenantId: category.GetTenantId(),
		Code:     category.GetCode(),
		Name:     category.GetName(),
		TransactionType: update.TransactionTypeInput{
			Code: category.GetTransactionType().GetCode(),
		},
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
	m.AssertNumberOfCalls(t, "Update", 1)
}
