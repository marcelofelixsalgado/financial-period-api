package find_test

import (
	"marcelofelixsalgado/financial-period-api/pkg/domain/period/entity"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/period/mocks"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/period/find"
	useCaseStatus "marcelofelixsalgado/financial-period-api/pkg/usecase/status"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFindPeriodUseCase_Execute(t *testing.T) {
	m := &mocks.PeriodRepositoryMock{}

	period, _ := entity.NewPeriod("1", "Period 1", "Period 1", 2023, time.Now(), time.Now(), time.Time{}, time.Time{})

	m.On("FindById", period.GetId()).Return(period, nil)

	useCase := find.NewFindUseCase(m)

	input := find.InputFindPeriodDto{
		Id: period.GetId(),
	}

	output, internalStatus, err := useCase.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.Id)
	assert.NotEmpty(t, output.CreatedAt)
	assert.Equal(t, period.GetName(), output.Name)
	assert.Equal(t, period.GetCode(), output.Code)
	assert.Equal(t, period.GetYear(), output.Year)
	assert.Equal(t, period.GetStartDate().Format(time.RFC3339), output.StartDate)
	assert.Equal(t, period.GetEndDate().Format(time.RFC3339), output.EndDate)
	assert.Equal(t, internalStatus, useCaseStatus.Success)
	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "FindById", 1)
}
