package find

import (
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/period"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/status"
	"time"
)

type IFindUseCase interface {
	Execute(InputFindPeriodDto) (OutputFindPeriodDto, status.InternalStatus, error)
}

type FindUseCase struct {
	repository period.IRepository
}

func NewFindUseCase(repository period.IRepository) IFindUseCase {
	return &FindUseCase{
		repository: repository,
	}
}

func (findUseCase *FindUseCase) Execute(input InputFindPeriodDto) (OutputFindPeriodDto, status.InternalStatus, error) {

	period, err := findUseCase.repository.Find(input.Id)
	if err != nil {
		return OutputFindPeriodDto{}, status.InternalServerError, err
	}
	if period == nil {
		return OutputFindPeriodDto{}, status.InvalidResourceId, err
	}

	if period.GetId() == "" {
		return OutputFindPeriodDto{}, status.NoRecordsFound, err
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

	return outputFindPeriodDto, status.Success, nil
}
