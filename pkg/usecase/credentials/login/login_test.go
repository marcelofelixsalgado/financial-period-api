package login_test

import (
	userCredentialsEntity "marcelofelixsalgado/financial-period-api/pkg/domain/credentials/entity"
	userEntity "marcelofelixsalgado/financial-period-api/pkg/domain/user/entity"
	userCredentialsRepositoryMock "marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/credentials/mocks"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/credentials/login"

	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	useCaseStatus "github.com/marcelofelixsalgado/financial-commons/pkg/usecase/status"
)

func TestLoginSucess(t *testing.T) {
	userCredentialsRepositoryMock := &userCredentialsRepositoryMock.UserCredentialsMock{}

	password := "pass"

	userEntity, _ := userEntity.Create("user1", "11", "111-1111", "user1@test.com")
	userCredentials, _ := userCredentialsEntity.Create(userEntity, password)

	userCredentialsRepositoryMock.On("FindByUserEmail", mock.Anything).Return(userCredentials, nil)

	useCase := login.NewLoginUseCase(userCredentialsRepositoryMock)

	input := login.InputUserLoginDto{
		Email:    userEntity.GetEmail(),
		Password: password,
	}

	output, internalStatus, err := useCase.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotNil(t, output.User.Id)
	assert.NotNil(t, output.User.Tenant.Id)
	assert.Equal(t, internalStatus, useCaseStatus.Success)
	userCredentialsRepositoryMock.AssertNumberOfCalls(t, "FindByUserEmail", 1)
	userCredentialsRepositoryMock.AssertExpectations(t)

}
