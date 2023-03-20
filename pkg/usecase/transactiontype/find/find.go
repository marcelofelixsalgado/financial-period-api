package find

import (
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/transactiontype"
	"marcelofelixsalgado/financial-period-api/pkg/usecase/status"
)

type IFindUseCase interface {
	Execute(InputFindTransactionTypeDto) (OutputFindTransactionTypeDto, status.InternalStatus, error)
}

type FindUseCase struct {
	repository transactiontype.ITransactionTypeRepository
}

func NewFindUseCase(repository transactiontype.ITransactionTypeRepository) IFindUseCase {
	return &FindUseCase{
		repository: repository,
	}
}

func (findUseCase *FindUseCase) Execute(input InputFindTransactionTypeDto) (OutputFindTransactionTypeDto, status.InternalStatus, error) {

	transactionType, err := findUseCase.repository.FindByCode(input.Code)
	if err != nil {
		return OutputFindTransactionTypeDto{}, status.InternalServerError, err
	}
	if transactionType == nil {
		return OutputFindTransactionTypeDto{}, status.InvalidResourceId, err
	}

	if transactionType.GetCode() == "" {
		return OutputFindTransactionTypeDto{}, status.NoRecordsFound, err
	}

	outputFindTransactionTypeDto := OutputFindTransactionTypeDto{
		Code: transactionType.GetCode(),
		Name: transactionType.GetName(),
	}

	return outputFindTransactionTypeDto, status.Success, nil
}
