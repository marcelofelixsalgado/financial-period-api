package update

import (
	"marcelofelixsalgado/financial-period-api/pkg/domain/period/entity"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/status"
	"time"
)

type IUpdateUseCase interface {
	Execute(InputUpdatePeriodDto) (OutputUpdatePeriodDto, status.InternalStatus, error)
}

type UpdateUseCase struct {
	repository repository.IRepository
}

func NewUpdateUseCase(repository repository.IRepository) IUpdateUseCase {
	return &UpdateUseCase{
		repository: repository,
	}
}

func (updateUseCase *UpdateUseCase) Execute(input InputUpdatePeriodDto) (OutputUpdatePeriodDto, status.InternalStatus, error) {

	startDate, err := time.Parse(time.RFC3339, input.StartDate)
	if err != nil {
		return OutputUpdatePeriodDto{}, status.InternalServerError, err
	}

	endDate, err := time.Parse(time.RFC3339, input.EndDate)
	if err != nil {
		return OutputUpdatePeriodDto{}, status.InternalServerError, err
	}

	// Find the entity before update
	currentEntity, err := updateUseCase.repository.Find(input.Id)
	if err != nil {
		return OutputUpdatePeriodDto{}, status.InternalServerError, err
	}
	if currentEntity == nil {
		return OutputUpdatePeriodDto{}, status.InvalidResourceId, err
	}

	entity, err := entity.NewPeriod(input.Id, input.Code, input.Name, input.Year, startDate, endDate, currentEntity.GetCreatedAt(), time.Now())
	if err != nil {
		return OutputUpdatePeriodDto{}, status.InternalServerError, err
	}

	// Persists in dabatase
	err = updateUseCase.repository.Update(entity)
	if err != nil {
		return OutputUpdatePeriodDto{}, status.InternalServerError, err
	}

	outputUpdatePeriodDto := OutputUpdatePeriodDto{
		Id:        entity.GetId(),
		Code:      entity.GetCode(),
		Name:      entity.GetName(),
		Year:      entity.GetYear(),
		StartDate: entity.GetStartDate().Format(time.RFC3339),
		EndDate:   entity.GetEndDate().Format(time.RFC3339),
		CreatedAt: entity.GetCreatedAt().Format(time.RFC3339),
		UpdatedAt: entity.GetUpdatedAt().Format(time.RFC3339),
	}

	return outputUpdatePeriodDto, status.Success, nil
}
