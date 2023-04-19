package delete

import (
	"github.com/marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/period"

	"github.com/marcelofelixsalgado/financial-commons/pkg/usecase/status"
)

type IDeleteUseCase interface {
	Execute(InputDeletePeriodDto) (OutputDeletePeriodDto, status.InternalStatus, error)
}

type DeleteUseCase struct {
	repository period.IPeriodRepository
}

func NewDeleteUseCase(repository period.IPeriodRepository) IDeleteUseCase {
	return &DeleteUseCase{
		repository: repository,
	}
}

func (deleteUseCase *DeleteUseCase) Execute(input InputDeletePeriodDto) (OutputDeletePeriodDto, status.InternalStatus, error) {

	// Find the entity before update
	entity, err := deleteUseCase.repository.FindById(input.Id)
	if err != nil {
		return OutputDeletePeriodDto{}, status.InternalServerError, err
	}
	if entity == nil {
		return OutputDeletePeriodDto{}, status.InvalidResourceId, nil
	}

	// Apply in dabatase
	err = deleteUseCase.repository.Delete(input.Id)
	if err != nil {
		return OutputDeletePeriodDto{}, status.InternalServerError, err
	}

	var outputDeletePeriodDto OutputDeletePeriodDto

	return outputDeletePeriodDto, status.Success, nil
}
