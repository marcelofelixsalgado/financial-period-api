package login

import (
	"marcelofelixsalgado/financial-period-api/pkg/domain/credentials/entity"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/auth"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/credentials"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/status"
)

type ILoginUseCase interface {
	Execute(InputUserLoginDto) (OutputUserLoginDto, status.InternalStatus, error)
}

type LoginUseCase struct {
	repository credentials.IUserCredentialsRepository
}

func NewLoginUseCase(repository credentials.IUserCredentialsRepository) ILoginUseCase {
	return &LoginUseCase{
		repository: repository,
	}
}

func (loginUseCase *LoginUseCase) Execute(input InputUserLoginDto) (OutputUserLoginDto, status.InternalStatus, error) {

	userCredentials, err := loginUseCase.repository.FindByUserEmail(input.Email)
	if err != nil {
		return OutputUserLoginDto{}, status.InternalServerError, err
	}
	if userCredentials == nil {
		return OutputUserLoginDto{}, status.InvalidResourceId, err
	}

	if userCredentials.GetId() == "" {
		return OutputUserLoginDto{}, status.NoRecordsFound, err
	}

	if err := entity.VerfifyPassword(userCredentials.GetPassword(), input.Password); err != nil {
		return OutputUserLoginDto{}, status.LoginFailed, err
	}

	accessToken, err := auth.CreateToken(userCredentials.GetUserId())
	if err != nil {
		return OutputUserLoginDto{}, status.InternalServerError, err
	}

	outputUserLoginDto := OutputUserLoginDto{
		User: userDto{
			Id: userCredentials.GetUserId(),
		},
		AccessToken: accessToken,
	}

	return outputUserLoginDto, status.Success, nil
}
