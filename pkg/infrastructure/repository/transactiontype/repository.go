package transactiontype

import (
	"github.com/marcelofelixsalgado/financial-period-api/pkg/domain/transactiontype/entity"
	"github.com/marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/filter"
)

type ITransactionTypeRepository interface {
	FindByCode(code string) (entity.ITransactionType, error)
	List(filterParameters []filter.FilterParameter) ([]entity.ITransactionType, error)
}
