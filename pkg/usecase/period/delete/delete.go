package delete

import (
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/status"
)

type IDeleteUseCase interface {
	Execute(InputDeletePeriodDto) (OutputDeletePeriodDto, status.InternalStatus, error)
}

type DeleteUseCase struct {
	repository repository.IRepository
}

func NewDeleteUseCase(repository repository.IRepository) IDeleteUseCase {
	return &DeleteUseCase{
		repository: repository,
	}
}

func (deleteUseCase *DeleteUseCase) Execute(input InputDeletePeriodDto) (OutputDeletePeriodDto, status.InternalStatus, error) {

	// Find the entity before update
	entity, err := deleteUseCase.repository.Find(input.Id)
	if err != nil {
		return OutputDeletePeriodDto{}, status.InternalServerError, err
	}
	if entity == nil {
		return OutputDeletePeriodDto{}, status.InvalidResourceId, err
	}

	// Apply in dabatase
	err = deleteUseCase.repository.Delete(input.Id)
	if err != nil {
		return OutputDeletePeriodDto{}, status.InternalServerError, err
	}

	var outputDeletePeriodDto OutputDeletePeriodDto

	return outputDeletePeriodDto, status.Success, nil
}
