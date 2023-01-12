package create_test

import (
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/period/mocks"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/period/create"
	useCaseStatus "marcelofelixsalgado/financial-period-api/pkg/usecase/status"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreatePeriodUseCase_Execute(t *testing.T) {
	m := &mocks.PeriodRepositoryMock{}
	m.On("Create", mock.Anything).Return(nil)

	useCase := create.NewCreateUseCase(m)

	input := create.InputCreatePeriodDto{
		Code:      "1",
		Name:      "Period 1",
		Year:      2023,
		StartDate: "2023-11-07T08:00:00Z",
		EndDate:   "2023-12-06T23:59:59Z",
	}

	output, internalStatus, err := useCase.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.Id)
	assert.NotEmpty(t, output.CreatedAt)
	assert.Equal(t, input.Code, output.Code)
	assert.Equal(t, input.Name, output.Name)
	assert.Equal(t, input.Year, output.Year)
	assert.Equal(t, input.StartDate, output.StartDate)
	assert.Equal(t, input.EndDate, output.EndDate)
	assert.Equal(t, internalStatus, useCaseStatus.Success)
	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "Create", 1)
}
