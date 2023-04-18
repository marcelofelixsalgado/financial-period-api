package create_test

import (
	tenantRepositoryMock "marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/tenant/mocks"
	userRepositoryMock "marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/user/mocks"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/user/create"

	useCaseStatus "github.com/marcelofelixsalgado/financial-commons/pkg/usecase/status"

	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateUserUseCase_Execute(t *testing.T) {
	userRepositoryMock := &userRepositoryMock.UserRepositoryMock{}
	tenantRepositoryMock := &tenantRepositoryMock.TenantRepositoryMock{}

	userRepositoryMock.On("Create", mock.Anything).Return(nil)
	tenantRepositoryMock.On("Create", mock.Anything).Return(nil)

	useCase := create.NewCreateUseCase(userRepositoryMock, tenantRepositoryMock)

	input := create.InputCreateUserDto{
		Name:  "user test",
		Phone: "123456",
		Email: "user@test.com",
	}

	output, internalStatus, err := useCase.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.Id)
	assert.NotEmpty(t, output.Tenant.Id)
	assert.NotEmpty(t, output.CreatedAt)
	assert.Equal(t, input.Name, output.Name)
	assert.Equal(t, input.Phone, output.Phone)
	assert.Equal(t, input.Email, output.Email)
	assert.Equal(t, internalStatus, useCaseStatus.Success)
	userRepositoryMock.AssertNumberOfCalls(t, "Create", 1)
	tenantRepositoryMock.AssertNumberOfCalls(t, "Create", 1)
	userRepositoryMock.AssertExpectations(t)
}
