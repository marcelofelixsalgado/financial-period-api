package list

import (
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository"
	"time"
)

func Execute(input InputListPeriodDto, repository repository.IRepository) (OutputListPeriodDto, error) {

	periods, err := repository.FindAll()
	if err != nil {
		return OutputListPeriodDto{}, err
	}

	outputListPeriodDto := OutputListPeriodDto{}

	for _, item := range periods {

		period := Period{
			Id:        item.GetId(),
			Code:      item.GetCode(),
			Name:      item.GetName(),
			Year:      item.GetYear(),
			StartDate: item.GetStartDate().Format(time.RFC3339),
			EndDate:   item.GetEndDate().Format(time.RFC3339),
		}

		outputListPeriodDto.Periods = append(outputListPeriodDto.Periods, period)
	}
	return outputListPeriodDto, nil
}
