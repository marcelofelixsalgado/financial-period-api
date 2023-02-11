package create_test

import (
	"errors"
	userCredentialsEntity "marcelofelixsalgado/financial-period-api/pkg/domain/credentials/entity"
	userEntity "marcelofelixsalgado/financial-period-api/pkg/domain/user/entity"

	userCredentialsRepositoryMock "marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/credentials/mocks"
	userRepositoryMock "marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/user/mocks"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/credentials/create"
	useCaseStatus "marcelofelixsalgado/financial-period-api/pkg/usecase/status"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateUserCredentialsSucess(t *testing.T) {
	userCredentialsRepositoryMock := &userCredentialsRepositoryMock.UserCredentialsMock{}
	userRepositoryMock := &userRepositoryMock.UserRepositoryMock{}

	userEntity, _ := userEntity.NewUser("123", "456", "user1", "111-1111", "user1@test.com", time.Time{}, time.Time{})

	userRepositoryMock.On("FindById", userEntity.GetId()).Return(userEntity, nil)
	userCredentialsRepositoryMock.On("FindByUserId", userEntity.GetId()).Return(nil, nil)
	userCredentialsRepositoryMock.On("Create", mock.Anything).Return(nil)

	useCase := create.NewCreateUseCase(userCredentialsRepositoryMock, userRepositoryMock)

	input := create.InputCreateUserCredentialsDto{
		UserId:   "123",
		Password: "pass",
	}

	output, internalStatus, err := useCase.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, internalStatus, useCaseStatus.Success)
	userRepositoryMock.AssertNumberOfCalls(t, "FindById", 1)
	userRepositoryMock.AssertNumberOfCalls(t, "FindByUserId", 0)
	userCredentialsRepositoryMock.AssertNumberOfCalls(t, "Create", 1)
	userCredentialsRepositoryMock.AssertExpectations(t)
}

func TestFindByIdError(t *testing.T) {
	userCredentialsRepositoryMock := &userCredentialsRepositoryMock.UserCredentialsMock{}
	userRepositoryMock := &userRepositoryMock.UserRepositoryMock{}

	userEntity, _ := userEntity.NewUser("123", "456", "user1", "111-1111", "user1@test.com", time.Time{}, time.Time{})

	userRepositoryMock.On("FindById", userEntity.GetId()).Return(nil, errors.New(""))

	useCase := create.NewCreateUseCase(userCredentialsRepositoryMock, userRepositoryMock)

	input := create.InputCreateUserCredentialsDto{
		UserId:   "123",
		Password: "pass",
	}

	output, _, err := useCase.Execute(input)

	assert.NotNil(t, err)
	assert.NotNil(t, output)
	userRepositoryMock.AssertNumberOfCalls(t, "FindById", 1)
	userRepositoryMock.AssertNumberOfCalls(t, "FindByUserId", 0)
	userCredentialsRepositoryMock.AssertNumberOfCalls(t, "Create", 0)
	userCredentialsRepositoryMock.AssertExpectations(t)
}

func TestFindByUserIdError(t *testing.T) {
	userCredentialsRepositoryMock := &userCredentialsRepositoryMock.UserCredentialsMock{}
	userRepositoryMock := &userRepositoryMock.UserRepositoryMock{}

	userEntity, _ := userEntity.NewUser("123", "456", "user1", "111-1111", "user1@test.com", time.Time{}, time.Time{})

	userRepositoryMock.On("FindById", userEntity.GetId()).Return(userEntity, nil)
	userCredentialsRepositoryMock.On("FindByUserId", userEntity.GetId()).Return(nil, errors.New(""))

	useCase := create.NewCreateUseCase(userCredentialsRepositoryMock, userRepositoryMock)

	input := create.InputCreateUserCredentialsDto{
		UserId:   "123",
		Password: "pass",
	}

	output, _, err := useCase.Execute(input)

	assert.NotNil(t, err)
	assert.NotNil(t, output)
	userRepositoryMock.AssertNumberOfCalls(t, "FindById", 1)
	userRepositoryMock.AssertNumberOfCalls(t, "FindByUserId", 0)
	userCredentialsRepositoryMock.AssertNumberOfCalls(t, "Create", 0)
	userCredentialsRepositoryMock.AssertExpectations(t)
}

func TestCreateError(t *testing.T) {
	userCredentialsRepositoryMock := &userCredentialsRepositoryMock.UserCredentialsMock{}
	userRepositoryMock := &userRepositoryMock.UserRepositoryMock{}

	userEntity, _ := userEntity.NewUser("123", "456", "user1", "111-1111", "user1@test.com", time.Time{}, time.Time{})

	userRepositoryMock.On("FindById", userEntity.GetId()).Return(userEntity, nil)
	userCredentialsRepositoryMock.On("FindByUserId", userEntity.GetId()).Return(nil, nil)
	userCredentialsRepositoryMock.On("Create", mock.Anything).Return(errors.New(""))

	useCase := create.NewCreateUseCase(userCredentialsRepositoryMock, userRepositoryMock)

	input := create.InputCreateUserCredentialsDto{
		UserId:   "123",
		Password: "pass",
	}

	output, _, err := useCase.Execute(input)

	assert.NotNil(t, err)
	assert.NotNil(t, output)
	userRepositoryMock.AssertNumberOfCalls(t, "FindById", 1)
	userRepositoryMock.AssertNumberOfCalls(t, "FindByUserId", 0)
	userCredentialsRepositoryMock.AssertNumberOfCalls(t, "Create", 1)
	userCredentialsRepositoryMock.AssertExpectations(t)
}

func TestUserCredentialsAreadyExists(t *testing.T) {
	userCredentialsRepositoryMock := &userCredentialsRepositoryMock.UserCredentialsMock{}
	userRepositoryMock := &userRepositoryMock.UserRepositoryMock{}

	userEntity, _ := userEntity.NewUser("123", "456", "user1", "111-1111", "user1@test.com", time.Time{}, time.Time{})
	userCredentials, _ := userCredentialsEntity.Create(userEntity, "pass")

	userRepositoryMock.On("FindById", userEntity.GetId()).Return(userEntity, nil)
	userCredentialsRepositoryMock.On("FindByUserId", userEntity.GetId()).Return(userCredentials, nil)

	useCase := create.NewCreateUseCase(userCredentialsRepositoryMock, userRepositoryMock)

	input := create.InputCreateUserCredentialsDto{
		UserId:   "123",
		Password: "pass",
	}

	output, internalStatus, err := useCase.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, internalStatus, useCaseStatus.EntityWithSameKeyAlreadyExists)
	userRepositoryMock.AssertNumberOfCalls(t, "FindById", 1)
	userRepositoryMock.AssertNumberOfCalls(t, "FindByUserId", 0)
	userCredentialsRepositoryMock.AssertNumberOfCalls(t, "Create", 0)
	userCredentialsRepositoryMock.AssertExpectations(t)
}
