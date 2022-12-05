package create

import (
	"marcelofelixsalgado/financial-period-api/pkg/domain/period/entity"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/period"

	"marcelofelixsalgado/financial-period-api/pkg/usecase/status"

	"time"
)

type ICreateUseCase interface {
	Execute(InputCreatePeriodDto) (OutputCreatePeriodDto, status.InternalStatus, error)
}

type CreateUseCase struct {
	repository period.IRepository
}

func NewCreateUseCase(repository period.IRepository) ICreateUseCase {
	return &CreateUseCase{
		repository: repository,
	}
}

func (createUseCase *CreateUseCase) Execute(input InputCreatePeriodDto) (OutputCreatePeriodDto, status.InternalStatus, error) {

	startDate, err := time.Parse(time.RFC3339, input.StartDate)
	if err != nil {
		return OutputCreatePeriodDto{}, status.InternalServerError, err
	}

	endDate, err := time.Parse(time.RFC3339, input.EndDate)
	if err != nil {
		return OutputCreatePeriodDto{}, status.InternalServerError, err
	}

	// Creates an entity
	entity, err := entity.Create(input.Code, input.Name, input.Year, startDate, endDate)
	if err != nil {
		return OutputCreatePeriodDto{}, status.InternalServerError, err
	}

	// Persists in dabatase
	err = createUseCase.repository.Create(entity)
	if err != nil {
		return OutputCreatePeriodDto{}, status.InternalServerError, err
	}

	outputCreatePeriodDto := OutputCreatePeriodDto{
		Id:        entity.GetId(),
		Code:      entity.GetCode(),
		Name:      entity.GetName(),
		Year:      entity.GetYear(),
		StartDate: entity.GetStartDate().Format(time.RFC3339),
		EndDate:   entity.GetEndDate().Format(time.RFC3339),
		CreatedAt: entity.GetCreatedAt().Format(time.RFC3339),
	}

	return outputCreatePeriodDto, status.Success, nil
}
