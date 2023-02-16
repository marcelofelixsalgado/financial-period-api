package balance

import (
	"marcelofelixsalgado/financial-period-api/pkg/domain/balance/entity"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/status"
)

type IBalanceRepository interface {
	Create(entity.IBalance) (status.RepositoryInternalStatus, error)
	Update(entity.IBalance) error
	FindById(id string) (entity.IBalance, error)
	List(tenantId string, periodId string) ([]entity.IBalance, error)
	Delete(id string) error
}
