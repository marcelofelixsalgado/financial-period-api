package create

import (
	"marcelofelixsalgado/financial-period-api/pkg/domain/period/entity"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository"

	"time"
)

func Execute(input InputCreatePeriodDto, repository repository.IRepository) (OutputCreatePeriodDto, error) {

	var outputCreatePeriodDto OutputCreatePeriodDto

	startDate, err := time.Parse(time.RFC3339, input.StartDate)
	if err != nil {
		return outputCreatePeriodDto, err
	}

	endDate, err := time.Parse(time.RFC3339, input.EndDate)
	if err != nil {
		return outputCreatePeriodDto, err
	}

	// Creates an entity
	entity, err := entity.Create(input.Code, input.Name, input.Year, startDate, endDate)
	if err != nil {
		return outputCreatePeriodDto, err
	}

	// Persists in dabatase
	err = repository.Create(entity)
	if err != nil {
		return outputCreatePeriodDto, err
	}

	outputCreatePeriodDto = OutputCreatePeriodDto{
		Id:        entity.GetId(),
		Code:      entity.GetCode(),
		Name:      entity.GetName(),
		Year:      entity.GetYear(),
		StartDate: entity.GetStartDate().Format(time.RFC3339),
		EndDate:   entity.GetEndDate().Format(time.RFC3339),
		CreatedAt: entity.GetCreatedAt().Format(time.RFC3339),
	}

	return outputCreatePeriodDto, nil
}
