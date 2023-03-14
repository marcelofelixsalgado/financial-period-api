package create_test

import (
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/group/mocks"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/group/create"
	useCaseStatus "marcelofelixsalgado/financial-period-api/pkg/usecase/status"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateGroupUseCase_Execute(t *testing.T) {
	m := &mocks.GroupRepositoryMock{}
	m.On("Create", mock.Anything).Return(nil)

	useCase := create.NewCreateUseCase(m)

	input := create.InputCreateGroupDto{
		TenantId: "123",
		Code:     "DF",
		Name:     "Despesa Fixa",
		Type: create.GroupType{
			Code: "EXP",
		},
	}

	output, internalStatus, err := useCase.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.Id)
	assert.NotEmpty(t, output.CreatedAt)
	assert.Equal(t, input.Code, output.Code)
	assert.Equal(t, input.Name, output.Name)
	assert.Equal(t, input.Type.Code, output.Type.Code)
	assert.Equal(t, internalStatus, useCaseStatus.Success)
	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "Create", 1)
}
