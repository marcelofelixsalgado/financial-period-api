package update

import (
	"marcelofelixsalgado/financial-period-api/pkg/domain/period/entity"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository"
	"time"
)

func Execute(input InputUpdatePeriodDto, repository repository.IRepository) (OutputUpdatePeriodDto, error) {
	var outputUpdatePeriodDto OutputUpdatePeriodDto

	startDate, err := time.Parse(time.RFC3339, input.StartDate)
	if err != nil {
		return outputUpdatePeriodDto, err
	}

	endDate, err := time.Parse(time.RFC3339, input.EndDate)
	if err != nil {
		return outputUpdatePeriodDto, err
	}

	// Find the entity before update
	findEntity, err := repository.FindById(input.Id)
	if err != nil {
		return outputUpdatePeriodDto, err
	}

	entity, err := entity.NewPeriod(input.Id, input.Code, input.Name, input.Year, startDate, endDate, findEntity.GetCreatedAt(), time.Now())
	if err != nil {
		return outputUpdatePeriodDto, err
	}

	// Persists in dabatase
	err = repository.Update(entity)
	if err != nil {
		return outputUpdatePeriodDto, err
	}

	outputUpdatePeriodDto = OutputUpdatePeriodDto{
		Id:        entity.GetId(),
		Code:      entity.GetCode(),
		Name:      entity.GetName(),
		Year:      entity.GetYear(),
		StartDate: entity.GetStartDate().Format(time.RFC3339),
		EndDate:   entity.GetEndDate().Format(time.RFC3339),
		CreatedAt: entity.GetCreatedAt().Format(time.RFC3339),
		UpdatedAt: entity.GetUpdatedAt().Format(time.RFC3339),
	}

	return outputUpdatePeriodDto, nil
}
