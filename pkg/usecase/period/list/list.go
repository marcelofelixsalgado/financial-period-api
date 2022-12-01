package list

import (
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/status"
	"time"
)

type IListUseCase interface {
	Execute(InputListPeriodDto, []repository.FilterParameter) (OutputListPeriodDto, status.InternalStatus, error)
}

type ListUseCase struct {
	repository repository.IRepository
}

func NewListUseCase(repository repository.IRepository) IListUseCase {
	return &ListUseCase{
		repository: repository,
	}
}

func (listUseCase *ListUseCase) Execute(input InputListPeriodDto, filterParameters []repository.FilterParameter) (OutputListPeriodDto, status.InternalStatus, error) {

	periods, err := listUseCase.repository.List(filterParameters)
	if err != nil {
		return OutputListPeriodDto{}, status.InternalServerError, err
	}

	outputListPeriodDto := OutputListPeriodDto{}

	if len(periods) == 0 {
		return outputListPeriodDto, status.NoRecordsFound, nil
	}

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
	return outputListPeriodDto, status.Success, nil
}
