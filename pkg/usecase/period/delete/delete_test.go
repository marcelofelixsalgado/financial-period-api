package delete_test

import (
	"marcelofelixsalgado/financial-period-api/pkg/domain/period/entity"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/period/mocks"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/period/delete"
	useCaseStatus "marcelofelixsalgado/financial-period-api/pkg/usecase/status"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDeletePeriodUseCase_Execute(t *testing.T) {
	m := &mocks.PeriodRepositoryMock{}

	period, _ := entity.NewPeriod("1", "Period 1", "Period 1", 2023, time.Now(), time.Now(), time.Time{}, time.Time{})

	m.On("FindById", period.GetId()).Return(period, nil)
	m.On("Delete", mock.Anything).Return(nil)

	useCase := delete.NewDeleteUseCase(m)

	input := delete.InputDeletePeriodDto{
		Id: period.GetId(),
	}

	output, internalStatus, err := useCase.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, internalStatus, useCaseStatus.Success)
	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "FindById", 1)
	m.AssertNumberOfCalls(t, "Delete", 1)
}