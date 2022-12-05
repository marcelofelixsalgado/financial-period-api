package login

import (
	"marcelofelixsalgado/financial-period-api/pkg/domain/user/entity"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/auth"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/user"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/status"
)

type ILoginUseCase interface {
	Execute(InputUserLoginDto) (OutputUserLoginDto, status.InternalStatus, error)
}

type LoginUseCase struct {
	repository user.IRepository
}

func NewLoginUseCase(repository user.IRepository) ILoginUseCase {
	return &LoginUseCase{
		repository: repository,
	}
}

func (loginUseCase *LoginUseCase) Execute(input InputUserLoginDto) (OutputUserLoginDto, status.InternalStatus, error) {

	user, err := loginUseCase.repository.FindByEmail(input.Email)
	if err != nil {
		return OutputUserLoginDto{}, status.InternalServerError, err
	}
	if user == nil {
		return OutputUserLoginDto{}, status.InvalidResourceId, err
	}

	if user.GetId() == "" {
		return OutputUserLoginDto{}, status.NoRecordsFound, err
	}

	if err := entity.VerfifyPassword(user.GetPassword(), input.Password); err != nil {
		return OutputUserLoginDto{}, status.LoginFailed, err
	}

	accessToken, err := auth.CreateToken(user.GetId())
	if err != nil {
		return OutputUserLoginDto{}, status.InternalServerError, err
	}

	outputUserLoginDto := OutputUserLoginDto{
		Email:       user.GetEmail(),
		AccessToken: accessToken,
	}

	return outputUserLoginDto, status.Success, nil
}
