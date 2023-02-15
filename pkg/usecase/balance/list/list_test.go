package list_test

import (
	"marcelofelixsalgado/financial-period-api/pkg/domain/balance/entity"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/balance/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"marcelofelixsalgado/financial-period-api/pkg/usecase/balance/list"
	useCaseStatus "marcelofelixsalgado/financial-period-api/pkg/usecase/status"
)

func TestListBalanceUseCase_Execute(t *testing.T) {
	m := &mocks.BalanceRepositoryMock{}

	tenantId := "111"
	periodId := ""

	balance1, _ := entity.NewBalance("1", tenantId, "11", "111", 100, 300, time.Now(), time.Now())
	balance2, _ := entity.NewBalance("2", tenantId, "22", "222", 200, 300, time.Now(), time.Now())

	balances := []entity.IBalance{balance1, balance2}

	m.On("List", tenantId, periodId).Return(balances, nil)

	useCase := list.NewListUseCase(m)

	input := list.InputListBalanceDto{
		TenantId: tenantId,
	}

	output, internalStatus, err := useCase.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Len(t, output.Balances, 2)

	assert.NotEmpty(t, output.Balances[0].Id)
	assert.Equal(t, balances[0].GetTenantId(), output.Balances[0].TenantId)
	assert.Equal(t, balances[0].GetPeriodId(), output.Balances[0].PeriodId)
	assert.Equal(t, balances[0].GetCategoryId(), output.Balances[0].CategoryId)
	assert.Equal(t, balances[0].GetActualAmount(), output.Balances[0].ActualAmount)
	assert.Equal(t, balances[0].GetLimitAmount(), output.Balances[0].LimitAmount)

	assert.Equal(t, balances[1].GetTenantId(), output.Balances[1].TenantId)
	assert.Equal(t, balances[1].GetPeriodId(), output.Balances[1].PeriodId)
	assert.Equal(t, balances[1].GetCategoryId(), output.Balances[1].CategoryId)
	assert.Equal(t, balances[1].GetActualAmount(), output.Balances[1].ActualAmount)
	assert.Equal(t, balances[1].GetLimitAmount(), output.Balances[1].LimitAmount)

	assert.Equal(t, internalStatus, useCaseStatus.Success)
	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "List", 1)
}
