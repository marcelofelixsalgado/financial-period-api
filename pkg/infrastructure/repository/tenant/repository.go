package tenant

import (
	"marcelofelixsalgado/financial-period-api/pkg/domain/tenant/entity"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/status"
)

type ITenantRepository interface {
	Create(entity.ITenant) (status.RepositoryInternalStatus, error)
}
