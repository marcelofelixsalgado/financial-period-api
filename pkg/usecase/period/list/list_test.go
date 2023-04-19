package list_test

import (
	"testing"
	"time"

	"github.com/marcelofelixsalgado/financial-period-api/pkg/domain/period/entity"
	"github.com/marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/filter"
	"github.com/marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/period/mocks"
	"github.com/marcelofelixsalgado/financial-period-api/pkg/usecase/period/list"

	useCaseStatus "github.com/marcelofelixsalgado/financial-commons/pkg/usecase/status"

	"github.com/stretchr/testify/assert"
)

func TestListPeriodUseCase_Execute(t *testing.T) {
	m := &mocks.PeriodRepositoryMock{}

	tenantId := "11"

	period1, _ := entity.NewPeriod("1", tenantId, "Period 1", "Period 1", 2023, time.Now(), time.Now(), time.Time{}, time.Time{})
	period2, _ := entity.NewPeriod("2", tenantId, "Period 2", "Period 2", 2024, time.Now(), time.Now(), time.Time{}, time.Time{})

	periods := []entity.IPeriod{period1, period2}

	m.On("List", []filter.FilterParameter{}, tenantId).Return(periods, nil)

	useCase := list.NewListUseCase(m)

	input := list.InputListPeriodDto{
		TenantId: tenantId,
	}

	output, internalStatus, err := useCase.Execute(input, []filter.FilterParameter{})

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Len(t, output.Periods, 2)

	assert.NotEmpty(t, output.Periods[0].Id)
	assert.Equal(t, periods[0].GetName(), output.Periods[0].Name)
	assert.Equal(t, periods[0].GetCode(), output.Periods[0].Code)
	assert.Equal(t, periods[0].GetYear(), output.Periods[0].Year)
	assert.Equal(t, periods[0].GetStartDate().Format(time.RFC3339), output.Periods[0].StartDate)
	assert.Equal(t, periods[0].GetEndDate().Format(time.RFC3339), output.Periods[0].EndDate)

	assert.Equal(t, periods[1].GetName(), output.Periods[1].Name)
	assert.Equal(t, periods[1].GetCode(), output.Periods[1].Code)
	assert.Equal(t, periods[1].GetYear(), output.Periods[1].Year)
	assert.Equal(t, periods[1].GetStartDate().Format(time.RFC3339), output.Periods[1].StartDate)
	assert.Equal(t, periods[1].GetEndDate().Format(time.RFC3339), output.Periods[1].EndDate)

	assert.Equal(t, internalStatus, useCaseStatus.Success)
	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "List", 1)
}
