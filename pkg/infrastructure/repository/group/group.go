package group

import (
	"database/sql"
	"errors"
	"marcelofelixsalgado/financial-period-api/pkg/domain/group/entity"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/filter"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/status"
	"time"

	"github.com/go-sql-driver/mysql"
)

type GroupRepository struct {
	client *sql.DB
}

type GroupModel struct {
	id        string
	tenantId  string
	code      string
	name      string
	groupType string
	createdAt time.Time
	updatedAt time.Time
}

func NewGroupRepository(client *sql.DB) IGroupRepository {
	return &GroupRepository{
		client: client,
	}
}

func (repository GroupRepository) Create(entity entity.IGroup) (status.RepositoryInternalStatus, error) {
	var mysqlErr *mysql.MySQLError

	model := GroupModel{
		id:        entity.GetId(),
		tenantId:  entity.GetTenantId(),
		code:      entity.GetCode(),
		name:      entity.GetName(),
		groupType: entity.GetGroupType().Code,
		createdAt: entity.GetCreatedAt(),
	}

	statement, err := repository.client.Prepare("insert into groups (id, tenant_id, code, name, type, created_at) values (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return status.InternalServerError, err
	}
	defer statement.Close()

	_, err = statement.Exec(model.id, model.tenantId, model.code, model.name, model.groupType, model.createdAt)
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		// Unique key violated
		return status.EntityWithSameKeyAlreadyExists, err
	}
	if err != nil {
		return status.InternalServerError, err
	}

	return status.Success, nil
}

func (repository GroupRepository) Update(entity entity.IGroup) (status.RepositoryInternalStatus, error) {
	var mysqlErr *mysql.MySQLError

	model := GroupModel{
		id:        entity.GetId(),
		tenantId:  entity.GetTenantId(),
		code:      entity.GetCode(),
		name:      entity.GetName(),
		groupType: entity.GetGroupType().Code,
		createdAt: entity.GetCreatedAt(),
		updatedAt: entity.GetUpdatedAt(),
	}

	statement, err := repository.client.Prepare("update groups set code = ?, name = ?, type = ?, updated_at = ? where id = ?")
	if err != nil {
		return status.InternalServerError, err
	}
	defer statement.Close()

	_, err = statement.Exec(model.code, model.name, model.groupType, model.updatedAt, model.id)
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		// Unique key violated
		return status.EntityWithSameKeyAlreadyExists, err
	}
	if err != nil {
		return status.InternalServerError, err
	}

	return status.Success, nil
}

func (repository GroupRepository) FindById(id string) (entity.IGroup, error) {

	row, err := repository.client.Query("select id, tenant_id, code, name, type, created_at, updated_at from groups where id = ?", id)
	if err != nil {
		return entity.Group{}, err
	}
	defer row.Close()

	var groupModel GroupModel
	if row.Next() {
		if err := row.Scan(&groupModel.id, &groupModel.tenantId, &groupModel.code, &groupModel.name, &groupModel.groupType, &groupModel.createdAt, &groupModel.updatedAt); err != nil {
			return entity.Group{}, err
		}

		groupType := entity.GroupType{
			Code: groupModel.groupType,
		}

		group, err := entity.NewGroup(groupModel.id, groupModel.tenantId, groupModel.code, groupModel.name, groupType, groupModel.createdAt, groupModel.updatedAt)
		if err != nil {
			return entity.Group{}, err
		}
		return group, nil
	}
	return nil, nil
}

func (repository GroupRepository) List(filterParameters []filter.FilterParameter, tenantId string) ([]entity.IGroup, error) {

	codeFilter := ""
	nameFilter := ""
	for _, filterParameter := range filterParameters {
		switch filterParameter.Name {
		case "code":
			codeFilter = filterParameter.Value
		case "name":
			nameFilter = filterParameter.Value
		}
	}

	var rows *sql.Rows
	var err error
	if len(filterParameters) == 0 {
		rows, err = repository.client.Query("select id, tenant_id, code, name, type, created_at, updated_at from groups where tenant_id = ?", tenantId)
	} else {
		if len(codeFilter) > 0 && len(nameFilter) == 0 {
			rows, err = repository.client.Query("select id, tenant_id, code, name, type, created_at, updated_at from groups where tenantId = ? and code = ?", tenantId, codeFilter)
		}
		if len(codeFilter) == 0 && len(nameFilter) > 0 {
			rows, err = repository.client.Query("select id, tenant_id, code, name, type, created_at, updated_at from groups where tenantId = ? and name = ?", tenantId, nameFilter)
		}
		if len(codeFilter) > 0 && len(nameFilter) > 0 {
			rows, err = repository.client.Query("select id, tenant_id, code, name, type, created_at, updated_at from groups where tenantId = ? and code = ? and name = ?", tenantId, codeFilter, nameFilter)
		}
	}

	if err != nil {
		return []entity.IGroup{}, err
	}
	defer rows.Close()

	groups := []entity.IGroup{}
	for rows.Next() {
		var groupModel GroupModel

		if err := rows.Scan(&groupModel.id, &groupModel.tenantId, &groupModel.code, &groupModel.name, &groupModel.groupType, &groupModel.createdAt, &groupModel.updatedAt); err != nil {
			return []entity.IGroup{}, err
		}

		groupType := entity.GroupType{
			Code: groupModel.groupType,
		}

		group, err := entity.NewGroup(groupModel.id, groupModel.tenantId, groupModel.code, groupModel.name, groupType, groupModel.createdAt, groupModel.updatedAt)
		if err != nil {
			return []entity.IGroup{}, err
		}

		groups = append(groups, group)
	}

	return groups, nil
}

func (repository GroupRepository) Delete(id string) error {

	statement, err := repository.client.Prepare("delete from groups where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
