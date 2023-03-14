package delete

import (
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/group"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/status"
)

type IDeleteUseCase interface {
	Execute(InputDeleteGroupDto) (OutputDeleteGroupDto, status.InternalStatus, error)
}

type DeleteUseCase struct {
	repository group.IGroupRepository
}

func NewDeleteUseCase(repository group.IGroupRepository) IDeleteUseCase {
	return &DeleteUseCase{
		repository: repository,
	}
}

func (deleteUseCase DeleteUseCase) Execute(input InputDeleteGroupDto) (OutputDeleteGroupDto, status.InternalStatus, error) {

	// Find the entity before update
	entity, err := deleteUseCase.repository.FindById(input.Id)
	if err != nil {
		return OutputDeleteGroupDto{}, status.InternalServerError, err
	}
	if entity == nil {
		return OutputDeleteGroupDto{}, status.InvalidResourceId, nil
	}

	// Apply in dabatase
	err = deleteUseCase.repository.Delete(input.Id)
	if err != nil {
		return OutputDeleteGroupDto{}, status.InternalServerError, err
	}

	var outputDeleteGroupDto OutputDeleteGroupDto

	return outputDeleteGroupDto, status.Success, nil
}
