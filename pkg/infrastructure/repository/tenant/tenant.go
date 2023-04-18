package tenant

import (
	"database/sql"
	"marcelofelixsalgado/financial-period-api/pkg/domain/tenant/entity"
	"time"

	"github.com/marcelofelixsalgado/financial-commons/pkg/infrastructure/repository/status"
)

type TenantRepository struct {
	client *sql.DB
}

type TenantModel struct {
	id        string
	createdAt time.Time
}

func NewTenantRepository(client *sql.DB) ITenantRepository {
	return &TenantRepository{
		client: client,
	}
}

func (repository *TenantRepository) Create(entity entity.ITenant) (status.RepositoryInternalStatus, error) {

	model := TenantModel{
		id:        entity.GetId(),
		createdAt: entity.GetCreatedAt(),
	}

	statement, err := repository.client.Prepare("insert into tenants (id, created_at) values (?, ?)")
	if err != nil {
		return status.InternalServerError, err
	}
	defer statement.Close()

	_, err = statement.Exec(model.id, model.createdAt)
	if err != nil {
		return status.InternalServerError, err
	}

	return status.Success, nil
}
