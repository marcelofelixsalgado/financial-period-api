package update_test

import (
	"testing"
	"time"

	"github.com/marcelofelixsalgado/financial-period-api/pkg/domain/balance/entity"
	"github.com/marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/balance/mocks"
	"github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/balance/update"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	useCaseStatus "github.com/marcelofelixsalgado/financial-commons/pkg/usecase/status"
)

func TestUpdateBalanceUseCase_Execute(t *testing.T) {
	m := &mocks.BalanceRepositoryMock{}

	balance, _ := entity.NewBalance("1", "11", "111", "1111", 100, 200, time.Now(), time.Now())

	m.On("FindById", balance.GetId()).Return(balance, nil)
	m.On("Update", mock.Anything).Return(nil)

	useCase := update.NewUpdateUseCase(m)

	input := update.InputUpdateBalanceDto{
		Id:           balance.GetId(),
		TenantId:     balance.GetTenantId(),
		ActualAmount: balance.GetActualAmount(),
	}

	output, internalStatus, err := useCase.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.Id)
	assert.NotEmpty(t, output.CreatedAt)
	assert.Equal(t, balance.GetActualAmount(), output.ActualAmount)
	assert.Equal(t, internalStatus, useCaseStatus.Success)
	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "FindById", 1)
	m.AssertNumberOfCalls(t, "Update", 1)
}
