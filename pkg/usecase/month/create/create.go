package create

import (
	"marcelofelixsalgado/financial-month-api/pkg/domain/month/entity"
	"marcelofelixsalgado/financial-month-api/pkg/infrastructure/repository"

	"time"
)

func Execute(input InputCreateMonthDto, repository repository.IRepository) (OutputCreateMonthDto, error) {

	var outputCreateMonthDto OutputCreateMonthDto

	startDate, err := time.Parse(time.RFC3339, input.StartDate)
	if err != nil {
		return outputCreateMonthDto, err
	}

	endDate, err := time.Parse(time.RFC3339, input.EndDate)
	if err != nil {
		return outputCreateMonthDto, err
	}

	// Creates an entity
	entity, err := entity.NewMonth(input.Code, input.Name, input.Year, startDate, endDate)
	if err != nil {
		return outputCreateMonthDto, err
	}

	// Persists in dabatase
	err = repository.Create(entity)
	if err != nil {
		return outputCreateMonthDto, err
	}

	outputCreateMonthDto = OutputCreateMonthDto{
		Id:        entity.GetId(),
		Code:      entity.GetCode(),
		Name:      entity.GetName(),
		Year:      entity.GetYear(),
		StartDate: entity.GetStartDate().String(),
		EndDate:   entity.GetEndDate().String(),
		CreatedAt: entity.GetCreatedAt(),
	}

	return outputCreateMonthDto, nil
}
