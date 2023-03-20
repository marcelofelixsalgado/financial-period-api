package delete_test

import (
	categoryEntity "marcelofelixsalgado/financial-period-api/pkg/domain/category/entity"
	"marcelofelixsalgado/financial-period-api/pkg/domain/subcategory/entity"
	transactionTypeEntity "marcelofelixsalgado/financial-period-api/pkg/domain/transactiontype/entity"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/subcategory/mocks"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/subcategory/delete"
	"testing"
	"time"

	useCaseStatus "marcelofelixsalgado/financial-period-api/pkg/usecase/status"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDeleteSubCategoryUseCase_Execute(t *testing.T) {
	m := &mocks.SubCategoryRepositoryMock{}

	tenantId := "11"
	transactionType, _ := transactionTypeEntity.NewTransactionType("EXPENSE", "Despesa")
	category, _ := categoryEntity.NewCategory("1", tenantId, "DF", "Despesa fixa", transactionType, time.Time{}, time.Time{})
	subCategory, _ := entity.NewSubCategory("111", tenantId, "CO", "Condimínio", category, time.Time{}, time.Time{})

	m.On("FindById", subCategory.GetId()).Return(subCategory, nil)
	m.On("Delete", mock.Anything).Return(nil)

	useCase := delete.NewDeleteUseCase(m)

	input := delete.InputDeleteUseCaseDto{
		Id: subCategory.GetId(),
	}

	output, internalStatus, err := useCase.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, internalStatus, useCaseStatus.Success)
	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "FindById", 1)
	m.AssertNumberOfCalls(t, "Delete", 1)
}
