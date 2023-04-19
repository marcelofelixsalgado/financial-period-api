package tenant

import (
	"github.com/marcelofelixsalgado/financial-period-api/pkg/domain/tenant/entity"

	"github.com/marcelofelixsalgado/financial-commons/pkg/infrastructure/repository/status"
)

type ITenantRepository interface {
	Create(entity.ITenant) (status.RepositoryInternalStatus, error)
}
