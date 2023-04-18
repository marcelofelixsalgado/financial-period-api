package list_test

import (
	"marcelofelixsalgado/financial-period-api/pkg/domain/transactiontype/entity"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/filter"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/transactiontype/mocks"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/transactiontype/list"
	"testing"

	useCaseStatus "github.com/marcelofelixsalgado/financial-commons/pkg/usecase/status"

	"github.com/stretchr/testify/assert"
)

func TestListTransactionTypeUseCase_Execute(t *testing.T) {
	m := &mocks.TransactionTypeRepositoryMock{}

	transactionType1, _ := entity.NewTransactionType("EXPENSE", "Despesa")
	transactionType2, _ := entity.NewTransactionType("EARNING", "Receita")

	transactionTypes := []entity.ITransactionType{transactionType1, transactionType2}

	m.On("List", []filter.FilterParameter{}).Return(transactionTypes, nil)

	useCase := list.NewListUseCase(m)

	input := list.InputListTransactionTypeDto{}

	output, internalStatus, err := useCase.Execute(input, []filter.FilterParameter{})

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Len(t, output.TransactionTypes, 2)

	assert.Equal(t, transactionTypes[0].GetName(), output.TransactionTypes[0].Name)
	assert.Equal(t, transactionTypes[0].GetCode(), output.TransactionTypes[0].Code)

	assert.Equal(t, transactionTypes[1].GetName(), output.TransactionTypes[1].Name)
	assert.Equal(t, transactionTypes[1].GetCode(), output.TransactionTypes[1].Code)

	assert.Equal(t, internalStatus, useCaseStatus.Success)
	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "List", 1)
}
