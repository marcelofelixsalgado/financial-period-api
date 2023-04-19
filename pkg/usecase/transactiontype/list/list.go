package list

import (
	"github.com/marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/filter"
	"github.com/marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/transactiontype"

	"github.com/marcelofelixsalgado/financial-commons/pkg/usecase/status"
)

type IListUseCase interface {
	Execute(InputListTransactionTypeDto, []filter.FilterParameter) (OutputListTransactionTypeDto, status.InternalStatus, error)
}

type ListUseCase struct {
	repository transactiontype.ITransactionTypeRepository
}

func NewListUseCase(repository transactiontype.ITransactionTypeRepository) IListUseCase {
	return &ListUseCase{
		repository: repository,
	}
}

func (listUseCase ListUseCase) Execute(input InputListTransactionTypeDto, filterParameters []filter.FilterParameter) (OutputListTransactionTypeDto, status.InternalStatus, error) {

	transactionTypes, err := listUseCase.repository.List(filterParameters)
	if err != nil {
		return OutputListTransactionTypeDto{}, status.InternalServerError, err
	}

	outputListTransactionTypeDto := OutputListTransactionTypeDto{}

	if len(transactionTypes) == 0 {
		return OutputListTransactionTypeDto{}, status.NoRecordsFound, nil
	}

	for _, item := range transactionTypes {
		transactionType := TransactionType{
			Code: item.GetCode(),
			Name: item.GetName(),
		}
		outputListTransactionTypeDto.TransactionTypes = append(outputListTransactionTypeDto.TransactionTypes, transactionType)
	}
	return outputListTransactionTypeDto, status.Success, nil
}
