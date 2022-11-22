package find

import "marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository"

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
		StartDate: period.GetStartDate().String(),
		EndDate:   period.GetEndDate().String(),
		CreatedAt: period.GetCreatedAt().String(),
	}

	if !period.GetUpdatedAt().IsZero() {
		outputFindPeriodDto.UpdatedAt = period.GetUpdatedAt().String()
	}

	return outputFindPeriodDto, nil
}
