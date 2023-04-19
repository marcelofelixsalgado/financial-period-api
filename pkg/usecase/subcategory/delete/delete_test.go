package delete_test

import (
	"testing"
	"time"

	categoryEntity "github.com/marcelofelixsalgado/financial-period-api/pkg/domain/category/entity"
	"github.com/marcelofelixsalgado/financial-period-api/pkg/domain/subcategory/entity"
	transactionTypeEntity "github.com/marcelofelixsalgado/financial-period-api/pkg/domain/transactiontype/entity"
	"github.com/marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/subcategory/mocks"
	"github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/subcategory/delete"

	useCaseStatus "github.com/marcelofelixsalgado/financial-commons/pkg/usecase/status"

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

	input := delete.InputDeleteSubCategoryDto{
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
