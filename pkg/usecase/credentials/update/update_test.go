package update_test

import (
	userCredentialsEntity "marcelofelixsalgado/financial-period-api/pkg/domain/credentials/entity"
	userEntity "marcelofelixsalgado/financial-period-api/pkg/domain/user/entity"
	userCredentialsRepositoryMock "marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/credentials/mocks"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/credentials/update"
	"time"

	useCaseStatus "github.com/marcelofelixsalgado/financial-commons/pkg/usecase/status"

	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdateSucess(t *testing.T) {
	userCredentialsRepositoryMock := &userCredentialsRepositoryMock.UserCredentialsMock{}

	password := "pass"

	userEntity, _ := userEntity.NewUser("123", "456", "user1", "111-1111", "user1@test.com", time.Time{}, time.Time{})
	userCredentials, _ := userCredentialsEntity.Create(userEntity, password)

	userCredentialsRepositoryMock.On("FindByUserId", userEntity.GetId()).Return(userCredentials, nil)
	userCredentialsRepositoryMock.On("Update", mock.Anything).Return(nil)

	useCase := update.NewUpdateUseCase(userCredentialsRepositoryMock)

	input := update.InputUpdateUserCredentialsDto{
		Id:              userCredentials.GetId(),
		UserId:          userEntity.GetId(),
		CurrentPassword: password,
		NewPassword:     "newpassword",
	}

	output, internalStatus, err := useCase.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, internalStatus, useCaseStatus.Success)
	assert.NotEmpty(t, output.Id)
	assert.NotEmpty(t, output.CreatedAt)
	assert.Equal(t, userCredentials.GetUserId(), output.UserId)
	assert.NotEqual(t, userCredentials.GetPassword(), output.Password)
	userCredentialsRepositoryMock.AssertNumberOfCalls(t, "FindByUserId", 1)
	userCredentialsRepositoryMock.AssertNumberOfCalls(t, "Update", 1)
	userCredentialsRepositoryMock.AssertExpectations(t)
}
