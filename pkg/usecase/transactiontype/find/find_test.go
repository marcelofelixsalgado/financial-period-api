package find_test

import (
	"testing"

	"github.com/marcelofelixsalgado/financial-period-api/pkg/domain/transactiontype/entity"
	"github.com/marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/transactiontype/mocks"
	"github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/transactiontype/find"

	useCaseStatus "github.com/marcelofelixsalgado/financial-commons/pkg/usecase/status"

	"github.com/stretchr/testify/assert"
)

func TestFindTransactionTypeUseCase_Execute(t *testing.T) {
	m := &mocks.TransactionTypeRepositoryMock{}

	transactionType, _ := entity.NewTransactionType("EXPENSE", "Expenses")

	m.On("FindByCode", transactionType.GetCode()).Return(transactionType, nil)

	useCase := find.NewFindUseCase(m)

	input := find.InputFindTransactionTypeDto{
		Code: transactionType.GetCode(),
	}

	output, internalStatus, err := useCase.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, transactionType.GetCode(), output.Code)
	assert.Equal(t, transactionType.GetName(), output.Name)
	assert.Equal(t, internalStatus, useCaseStatus.Success)
	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "FindByCode", 1)
}
