package delete

import (
	"github.com/marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/user"

	"github.com/marcelofelixsalgado/financial-commons/pkg/usecase/status"
)

type IDeleteUseCase interface {
	Execute(InputDeleteUserDto) (OutputDeleteUserDto, status.InternalStatus, error)
}

type DeleteUseCase struct {
	repository user.IUserRepository
}

func NewDeleteUseCase(repository user.IUserRepository) IDeleteUseCase {
	return &DeleteUseCase{
		repository: repository,
	}
}

func (deleteUseCase *DeleteUseCase) Execute(input InputDeleteUserDto) (OutputDeleteUserDto, status.InternalStatus, error) {

	// Find the entity before update
	entity, err := deleteUseCase.repository.FindById(input.Id)
	if err != nil {
		return OutputDeleteUserDto{}, status.InternalServerError, err
	}
	if entity == nil {
		return OutputDeleteUserDto{}, status.InvalidResourceId, err
	}

	// Apply in dabatase
	err = deleteUseCase.repository.Delete(input.Id)
	if err != nil {
		return OutputDeleteUserDto{}, status.InternalServerError, err
	}

	var outputDeleteUserDto OutputDeleteUserDto

	return outputDeleteUserDto, status.Success, nil
}
