package list

import (
	"time"

	"github.com/marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/filter"
	"github.com/marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/period"

	"github.com/marcelofelixsalgado/financial-commons/pkg/usecase/status"
)

type IListUseCase interface {
	Execute(InputListPeriodDto, []filter.FilterParameter) (OutputListPeriodDto, status.InternalStatus, error)
}

type ListUseCase struct {
	repository period.IPeriodRepository
}

func NewListUseCase(repository period.IPeriodRepository) IListUseCase {
	return &ListUseCase{
		repository: repository,
	}
}

func (listUseCase *ListUseCase) Execute(input InputListPeriodDto, filterParameters []filter.FilterParameter) (OutputListPeriodDto, status.InternalStatus, error) {

	periods, err := listUseCase.repository.List(filterParameters, input.TenantId)
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
