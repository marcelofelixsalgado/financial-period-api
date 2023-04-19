package delete

import (
	"github.com/marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/balance"

	"github.com/marcelofelixsalgado/financial-commons/pkg/usecase/status"
)

type IDeleteUseCase interface {
	Execute(InputDeleteBalanceDto) (OutputDeleteBalanceDto, status.InternalStatus, error)
}

type DeleteUseCase struct {
	repository balance.IBalanceRepository
}

func NewDeleteUseCase(repository balance.IBalanceRepository) IDeleteUseCase {
	return &DeleteUseCase{
		repository: repository,
	}
}

func (deleteUseCase *DeleteUseCase) Execute(input InputDeleteBalanceDto) (OutputDeleteBalanceDto, status.InternalStatus, error) {

	entity, err := deleteUseCase.repository.FindById(input.Id)
	if err != nil {
		return OutputDeleteBalanceDto{}, status.InternalServerError, err
	}
	if entity == nil {
		return OutputDeleteBalanceDto{}, status.InvalidResourceId, nil
	}

	// Apply in database
	err = deleteUseCase.repository.Delete(input.Id)
	if err != nil {
		return OutputDeleteBalanceDto{}, status.InternalServerError, err
	}

	var outputDeleteBalanceDto OutputDeleteBalanceDto

	return outputDeleteBalanceDto, status.Success, nil
}
