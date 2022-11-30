package delete

import (
	"marcelofelixsalgado/financial-period-api/pkg/domain/period/entity"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository"
)

type IDeleteUseCase interface {
	Execute(InputDeletePeriodDto) (OutputDeletePeriodDto, error)
}

type DeleteUseCase struct {
	repository repository.IRepository
}

func NewDeleteUseCase(repository repository.IRepository) IDeleteUseCase {
	return &DeleteUseCase{
		repository: repository,
	}
}

func (deleteUseCase *DeleteUseCase) Execute(input InputDeletePeriodDto) (OutputDeletePeriodDto, error) {

	var outputDeletePeriodDto OutputDeletePeriodDto

	// Find the entity before update
	findEntity, err := deleteUseCase.repository.FindById(input.Id)
	if err != nil {
		return outputDeletePeriodDto, err
	}

	_, err = entity.NewPeriod(input.Id, findEntity.GetCode(), findEntity.GetName(), findEntity.GetYear(), findEntity.GetStartDate(), findEntity.GetEndDate(), findEntity.GetCreatedAt(), findEntity.GetUpdatedAt())
	if err != nil {
		return outputDeletePeriodDto, err
	}

	// Apply in dabatase
	err = deleteUseCase.repository.Delete(input.Id)
	if err != nil {
		return outputDeletePeriodDto, err
	}

	return outputDeletePeriodDto, nil
}
