package update

import (
	"marcelofelixsalgado/financial-period-api/pkg/domain/balance/entity"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/balance"
	"time"

	"github.com/marcelofelixsalgado/financial-commons/pkg/usecase/status"
)

type IUpdateUseCase interface {
	Execute(InputUpdateBalanceDto) (OutputUpdateBalanceDto, status.InternalStatus, error)
}

type UpdateUseCase struct {
	repository balance.IBalanceRepository
}

func NewUpdateUseCase(repository balance.IBalanceRepository) IUpdateUseCase {
	return &UpdateUseCase{
		repository: repository,
	}
}

func (updateUseCase *UpdateUseCase) Execute(input InputUpdateBalanceDto) (OutputUpdateBalanceDto, status.InternalStatus, error) {

	// Find the entity before update
	currentEntity, err := updateUseCase.repository.FindById(input.Id)
	if err != nil {
		return OutputUpdateBalanceDto{}, status.InternalServerError, err
	}
	if currentEntity == nil {
		return OutputUpdateBalanceDto{}, status.InvalidResourceId, nil
	}

	entity, err := entity.NewBalance(input.Id, input.TenantId, currentEntity.GetPeriodId(), currentEntity.GetCategoryId(), input.ActualAmount, currentEntity.GetLimitAmount(), currentEntity.GetCreatedAt(), time.Now())
	if err != nil {
		return OutputUpdateBalanceDto{}, status.InternalServerError, err
	}

	// Persists in database
	err = updateUseCase.repository.Update(entity)
	if err != nil {
		return OutputUpdateBalanceDto{}, status.InternalServerError, err
	}

	outputUpdateBalanceDto := OutputUpdateBalanceDto{
		Id:           entity.GetId(),
		PeriodId:     entity.GetPeriodId(),
		CategoryId:   entity.GetCategoryId(),
		ActualAmount: entity.GetActualAmount(),
		LimitAmount:  entity.GetLimitAmount(),
		CreatedAt:    entity.GetCreatedAt().Format(time.RFC3339),
		UpdatedAt:    entity.GetUpdatedAt().Format(time.RFC3339),
	}

	return outputUpdateBalanceDto, status.Success, nil
}
