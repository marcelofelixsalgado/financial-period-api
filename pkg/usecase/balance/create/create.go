package create

import (
	"marcelofelixsalgado/financial-period-api/pkg/domain/balance/entity"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/balance"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/status"

	repositoryStatus "marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/status"
)

type ICreateUseCase interface {
	Execute(InputCreateBalanceDto) (OutputCreateBalanceDto, status.InternalStatus, error)
}

type CreateUseCase struct {
	repository balance.IBalanceRepository
}

func NewCreateUseCase(repository balance.IBalanceRepository) ICreateUseCase {
	return &CreateUseCase{
		repository: repository,
	}
}

func (createUseCase *CreateUseCase) Execute(input InputCreateBalanceDto) (OutputCreateBalanceDto, status.InternalStatus, error) {

	// Creates an entity
	entity, err := entity.Create(input.TenantId, input.PeriodId, input.CategoryId, input.ActualAmount, input.LimitAmount)
	if err != nil {
		return OutputCreateBalanceDto{}, status.InternalServerError, err
	}

	// Persists in database
	repositoryInternalStatus, err := createUseCase.repository.Create(entity)
	if repositoryInternalStatus == repositoryStatus.EntityWithSameKeyAlreadyExists {
		return OutputCreateBalanceDto{}, status.EntityWithSameKeyAlreadyExists, err
	}
	if err != nil {
		return OutputCreateBalanceDto{}, status.InternalServerError, err
	}

	outputCreateBalanceDto := OutputCreateBalanceDto{
		Id:           entity.GetId(),
		PeriodId:     entity.GetPeriodId(),
		CategoryId:   entity.GetCategoryId(),
		ActualAmount: entity.GetActualAmount(),
		LimitAmount:  entity.GetLimitAmount(),
		CreatedAt:    entity.GetCreatedAt(),
	}

	return outputCreateBalanceDto, status.Success, nil
}
