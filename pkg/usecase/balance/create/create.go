package create

import (
	"marcelofelixsalgado/financial-period-api/pkg/domain/balance/entity"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/balance"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/status"
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
	err = createUseCase.repository.Create(entity)
	if err != nil {
		return OutputCreateBalanceDto{}, status.InternalServerError, err
	}

	outputCreateBalanceDto := OutputCreateBalanceDto{
		Id:           entity.GetId(),
		TenantId:     entity.GetTenantId(),
		PeriodId:     entity.GetPeriodId(),
		CategoryId:   entity.GetCategoryId(),
		ActualAmount: entity.GetActualAmount(),
		LimitAmount:  entity.GetLimitAmount(),
		CreatedAt:    entity.GetCreatedAt(),
	}

	return outputCreateBalanceDto, status.Success, nil
}
