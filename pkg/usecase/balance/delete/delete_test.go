package delete_test

import (
	"testing"
	"time"

	"github.com/marcelofelixsalgado/financial-period-api/pkg/domain/balance/entity"
	"github.com/marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/balance/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/balance/delete"

	useCaseStatus "github.com/marcelofelixsalgado/financial-commons/pkg/usecase/status"
)

func TestDeleteBalanceUseCase_Execute(t *testing.T) {
	m := &mocks.BalanceRepositoryMock{}

	balance, _ := entity.NewBalance("123", "456", "789", "012", 100, 200, time.Now(), time.Now())

	m.On("FindById", balance.GetId()).Return(balance, nil)
	m.On("Delete", mock.Anything).Return(nil)

	useCase := delete.NewDeleteUseCase(m)

	input := delete.InputDeleteBalanceDto{
		Id: balance.GetId(),
	}

	output, internalStatus, err := useCase.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, internalStatus, useCaseStatus.Success)
	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "FindById", 1)
	m.AssertNumberOfCalls(t, "Delete", 1)
}
