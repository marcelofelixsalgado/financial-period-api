package find

import (
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/balance"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/status"
	"time"
)

type IFindUseCase interface {
	Execute(InputFindBalanceDto) (OutputFindBalanceDto, status.InternalStatus, error)
}

type FindUseCase struct {
	repository balance.IBalanceRepository
}

func NewFindUseCase(repository balance.IBalanceRepository) IFindUseCase {
	return &FindUseCase{
		repository: repository,
	}
}

func (findUseCase *FindUseCase) Execute(input InputFindBalanceDto) (OutputFindBalanceDto, status.InternalStatus, error) {

	balance, err := findUseCase.repository.FindById(input.Id)
	if err != nil {
		return OutputFindBalanceDto{}, status.InternalServerError, err
	}
	if balance == nil {
		return OutputFindBalanceDto{}, status.InvalidResourceId, err
	}
	if balance.GetId() == "" {
		return OutputFindBalanceDto{}, status.NoRecordsFound, err
	}

	outputFindBalanceDto := OutputFindBalanceDto{
		Id:           balance.GetId(),
		TenantId:     balance.GetTenantId(),
		PeriodId:     balance.GetPeriodId(),
		CategoryId:   balance.GetCategoryId(),
		ActualAmount: balance.GetActualAmount(),
		LimitAmount:  balance.GetLimitAmount(),
		CreatedAt:    balance.GetCreatedAt().Format(time.RFC3339),
		UpdatedAt:    balance.GetUpdatedAt().Format(time.RFC3339),
	}

	return outputFindBalanceDto, status.Success, nil
}
