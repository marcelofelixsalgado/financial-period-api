package delete

import (
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/user"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/status"
)

type IDeleteUseCase interface {
	Execute(InputDeleteUserDto) (OutputDeleteUserDto, status.InternalStatus, error)
}

type DeleteUseCase struct {
	repository user.IRepository
}

func NewDeleteUseCase(repository user.IRepository) IDeleteUseCase {
	return &DeleteUseCase{
		repository: repository,
	}
}

func (deleteUseCase *DeleteUseCase) Execute(input InputDeleteUserDto) (OutputDeleteUserDto, status.InternalStatus, error) {

	// Find the entity before update
	entity, err := deleteUseCase.repository.Find(input.Id)
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
