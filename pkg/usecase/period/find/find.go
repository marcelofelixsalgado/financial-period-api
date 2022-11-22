package find

import (
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository"
	"time"
)

func Execute(input InputFindPeriodDto, repository repository.IRepository) (OutputFindPeriodDto, error) {

	period, err := repository.FindById(input.Id)
	if err != nil {
		return OutputFindPeriodDto{}, err
	}

	outputFindPeriodDto := OutputFindPeriodDto{
		Id:        period.GetId(),
		Code:      period.GetCode(),
		Name:      period.GetName(),
		Year:      period.GetYear(),
		StartDate: period.GetStartDate().Format(time.RFC3339),
		EndDate:   period.GetEndDate().Format(time.RFC3339),
		CreatedAt: period.GetCreatedAt().Format(time.RFC3339),
	}

	if !period.GetUpdatedAt().IsZero() {
		outputFindPeriodDto.UpdatedAt = period.GetUpdatedAt().Format(time.RFC3339)
	}

	return outputFindPeriodDto, nil
}