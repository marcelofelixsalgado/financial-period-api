package find_test

import (
	"marcelofelixsalgado/financial-period-api/pkg/domain/balance/entity"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/balance/mocks"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/balance/find"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	useCaseStatus "github.com/marcelofelixsalgado/financial-commons/pkg/usecase/status"
)

func TestFindBalanceUseCase_Execute(t *testing.T) {
	m := &mocks.BalanceRepositoryMock{}

	balance, _ := entity.NewBalance("123", "456", "789", "012", 100, 200, time.Now(), time.Now())

	m.On("FindById", balance.GetId()).Return(balance, nil)

	useCase := find.NewFindUseCase(m)

	input := find.InputFindBalanceDto{
		Id: balance.GetId(),
	}

	output, internalStatus, err := useCase.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.Id)
	assert.NotEmpty(t, output.CreatedAt)
	assert.Equal(t, balance.GetPeriodId(), output.PeriodId)
	assert.Equal(t, balance.GetCategoryId(), output.CategoryId)
	assert.Equal(t, balance.GetActualAmount(), output.ActualAmount)
	assert.Equal(t, balance.GetLimitAmount(), output.LimitAmount)
	assert.Equal(t, internalStatus, useCaseStatus.Success)
	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "FindById", 1)
}
