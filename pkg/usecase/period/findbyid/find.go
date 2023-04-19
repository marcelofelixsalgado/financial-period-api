package findbyid

import (
	"time"

	"github.com/marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/period"

	"github.com/marcelofelixsalgado/financial-commons/pkg/usecase/status"
)

type IFindByIdUseCase interface {
	Execute(InputFindByIdPeriodDto) (OutputFindByIdPeriodDto, status.InternalStatus, error)
}

type FindByIdUseCase struct {
	repository period.IPeriodRepository
}

func NewFindByIdUseCase(repository period.IPeriodRepository) IFindByIdUseCase {
	return &FindByIdUseCase{
		repository: repository,
	}
}

func (findByIdUseCase *FindByIdUseCase) Execute(input InputFindByIdPeriodDto) (OutputFindByIdPeriodDto, status.InternalStatus, error) {

	period, err := findByIdUseCase.repository.FindById(input.Id)
	if err != nil {
		return OutputFindByIdPeriodDto{}, status.InternalServerError, err
	}
	if period == nil {
		return OutputFindByIdPeriodDto{}, status.InvalidResourceId, err
	}

	if period.GetId() == "" {
		return OutputFindByIdPeriodDto{}, status.NoRecordsFound, err
	}

	outputFindPeriodDto := OutputFindByIdPeriodDto{
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
