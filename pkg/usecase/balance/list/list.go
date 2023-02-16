package list

import (
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/balance"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/status"
)

type IListUseCase interface {
	Execute(InputListBalanceDto) (OutputListBalanceDto, status.InternalStatus, error)
}

type ListUseCase struct {
	repository balance.IBalanceRepository
}

func NewListUseCase(repository balance.IBalanceRepository) IListUseCase {
	return &ListUseCase{
		repository: repository,
	}
}

func (listUseCase *ListUseCase) Execute(input InputListBalanceDto) (OutputListBalanceDto, status.InternalStatus, error) {

	balances, err := listUseCase.repository.List(input.TenantId, input.PeriodId)
	if err != nil {
		return OutputListBalanceDto{}, status.InternalServerError, err
	}

	outputListBalanceDto := OutputListBalanceDto{}

	if len(balances) == 0 {
		return outputListBalanceDto, status.NoRecordsFound, nil
	}

	for _, item := range balances {
		balance := Balance{
			Id:           item.GetId(),
			PeriodId:     item.GetPeriodId(),
			CategoryId:   item.GetCategoryId(),
			ActualAmount: item.GetActualAmount(),
			LimitAmount:  item.GetLimitAmount(),
		}
		outputListBalanceDto.Balances = append(outputListBalanceDto.Balances, balance)
	}

	return outputListBalanceDto, status.Success, nil
}
