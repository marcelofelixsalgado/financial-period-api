package delete

import (
	"marcelofelixsalgado/financial-period-api/pkg/domain/period/entity"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository"
)

func Execute(input InputDeletePeriodDto, repository repository.IRepository) (OutputDeletePeriodDto, error) {

	var outputDeletePeriodDto OutputDeletePeriodDto

	// Find the entity before update
	findEntity, err := repository.FindById(input.Id)
	if err != nil {
		return outputDeletePeriodDto, err
	}

	_, err = entity.NewPeriod(input.Id, findEntity.GetCode(), findEntity.GetName(), findEntity.GetYear(), findEntity.GetStartDate(), findEntity.GetEndDate(), findEntity.GetCreatedAt(), findEntity.GetUpdatedAt())
	if err != nil {
		return outputDeletePeriodDto, err
	}

	// Apply in dabatase
	err = repository.Delete(input.Id)
	if err != nil {
		return outputDeletePeriodDto, err
	}

	return outputDeletePeriodDto, nil
}
