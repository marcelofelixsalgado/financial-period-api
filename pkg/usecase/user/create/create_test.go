package create_test

import (
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/user/mocks"
	useCaseStatus "marcelofelixsalgado/financial-period-api/pkg/usecase/status"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/user/create"

	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateUserUseCase_Execute(t *testing.T) {
	m := &mocks.UserRepositoryMock{}
	m.On("Create", mock.Anything).Return(nil)

	useCase := create.NewCreateUseCase(m)

	input := create.InputCreateUserDto{
		Name:  "user test",
		Phone: "123456",
		Email: "user@test.com",
	}

	output, internalStatus, err := useCase.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.Id)
	assert.NotEmpty(t, output.CreatedAt)
	assert.Equal(t, input.Name, output.Name)
	assert.Equal(t, input.Phone, output.Phone)
	assert.Equal(t, input.Email, output.Email)
	assert.Equal(t, internalStatus, useCaseStatus.Success)
	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "Create", 1)
}
