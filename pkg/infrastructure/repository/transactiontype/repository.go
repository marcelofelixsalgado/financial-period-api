package transactiontype

import (
	"marcelofelixsalgado/financial-period-api/pkg/domain/transactiontype/entity"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/filter"
)

type ITransactionTypeRepository interface {
	FindByCode(code string) (entity.ITransactionType, error)
	List(filterParameters []filter.FilterParameter) ([]entity.ITransactionType, error)
}
