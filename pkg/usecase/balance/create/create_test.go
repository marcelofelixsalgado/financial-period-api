package create_test

import (
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/balance/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"marcelofelixsalgado/financial-period-api/pkg/usecase/balance/create"
	useCaseStatus "marcelofelixsalgado/financial-period-api/pkg/usecase/status"
)

func TestCreateBalanceUseCase_Execute(t *testing.T) {
	m := &mocks.BalanceRepositoryMock{}
	m.On("Create", mock.Anything).Return(nil)

	useCase := create.NewCreateUseCase(m)

	input := create.InputCreateBalanceDto{
		TenantId:     "123",
		PeriodId:     "456",
		CategoryId:   "789",
		ActualAmount: 104.35,
		LimitAmount:  200,
	}

	output, internalStatus, err := useCase.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.Id)
	assert.NotEmpty(t, output.CreatedAt)
	assert.Equal(t, input.PeriodId, output.PeriodId)
	assert.Equal(t, input.CategoryId, output.CategoryId)
	assert.Equal(t, input.ActualAmount, output.ActualAmount)
	assert.Equal(t, input.LimitAmount, output.LimitAmount)
	assert.Equal(t, internalStatus, useCaseStatus.Success)
	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "Create", 1)
}
