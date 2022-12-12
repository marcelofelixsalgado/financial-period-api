package create

import (
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/user"

	"marcelofelixsalgado/financial-period-api/pkg/domain/credentials/entity"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/credentials"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/status"

	"time"
)

type ICreateUseCase interface {
	Execute(InputCreateUserCredentialsDto) (OutputCreateUserCredentialsDto, status.InternalStatus, error)
}

type CreateUseCase struct {
	userCredentialsRepository credentials.IUserCredentialsRepository
	usersRepository           user.IUserRepository
}

func NewCreateUseCase(userCredentialsRepository credentials.IUserCredentialsRepository, usersRepository user.IUserRepository) ICreateUseCase {
	return &CreateUseCase{
		userCredentialsRepository: userCredentialsRepository,
		usersRepository:           usersRepository,
	}
}

func (createUseCase *CreateUseCase) Execute(input InputCreateUserCredentialsDto) (OutputCreateUserCredentialsDto, status.InternalStatus, error) {

	user, err := createUseCase.usersRepository.FindById(input.UserId)
	if err != nil {
		return OutputCreateUserCredentialsDto{}, status.InternalServerError, err
	}

	// Creates an entity
	userCredentials, err := entity.Create(user, input.Password)
	if err != nil {
		return OutputCreateUserCredentialsDto{}, status.InternalServerError, err
	}

	// Persists the user
	err = createUseCase.userCredentialsRepository.Create(userCredentials)
	if err != nil {
		return OutputCreateUserCredentialsDto{}, status.InternalServerError, err
	}

	outputCreateUserCredentialsDto := OutputCreateUserCredentialsDto{
		CreatedAt: user.GetCreatedAt().Format(time.RFC3339),
	}

	return outputCreateUserCredentialsDto, status.Success, nil
}
