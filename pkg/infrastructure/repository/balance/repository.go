package balance

import (
	"marcelofelixsalgado/financial-period-api/pkg/domain/balance/entity"
)

type IBalanceRepository interface {
	Create(entity.IBalance) error
	Update(entity.IBalance) error
	FindById(id string) (entity.IBalance, error)
	List(tenantId string, periodId string) ([]entity.IBalance, error)
	Delete(id string) error
}
