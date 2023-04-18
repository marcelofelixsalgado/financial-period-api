package category

import (
	"database/sql"
	"errors"
	"marcelofelixsalgado/financial-period-api/pkg/domain/category/entity"
	transactionTypeEntity "marcelofelixsalgado/financial-period-api/pkg/domain/transactiontype/entity"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/filter"
	transactionTypeModel "marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/transactiontype"
	"time"

	"github.com/marcelofelixsalgado/financial-commons/pkg/infrastructure/repository/status"

	"github.com/go-sql-driver/mysql"
)

type CategoryRepository struct {
	client *sql.DB
}

type CategoryModel struct {
	Id              string
	TenantId        string
	Code            string
	Name            string
	TransactionType transactionTypeModel.TransactionTypeModel
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func NewCategoryRepository(client *sql.DB) ICategoryRepository {
	return &CategoryRepository{
		client: client,
	}
}

func (repository CategoryRepository) Create(entity entity.ICategory) (status.RepositoryInternalStatus, error) {
	var mysqlErr *mysql.MySQLError

	model := CategoryModel{
		Id:       entity.GetId(),
		TenantId: entity.GetTenantId(),
		Code:     entity.GetCode(),
		Name:     entity.GetName(),
		TransactionType: transactionTypeModel.TransactionTypeModel{
			Code: entity.GetTransactionType().GetCode(),
		},
		CreatedAt: entity.GetCreatedAt(),
	}

	statement, err := repository.client.Prepare("insert into categories (id, tenant_id, code, name, transaction_type_code, created_at) values (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return status.InternalServerError, err
	}
	defer statement.Close()

	_, err = statement.Exec(model.Id, model.TenantId, model.Code, model.Name, model.TransactionType.Code, model.CreatedAt)
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		// Unique key violated
		return status.EntityWithSameKeyAlreadyExists, err
	}
	if err != nil {
		return status.InternalServerError, err
	}

	return status.Success, nil
}

func (repository CategoryRepository) Update(entity entity.ICategory) (status.RepositoryInternalStatus, error) {
	var mysqlErr *mysql.MySQLError

	model := CategoryModel{
		Id:       entity.GetId(),
		TenantId: entity.GetTenantId(),
		Code:     entity.GetCode(),
		Name:     entity.GetName(),
		TransactionType: transactionTypeModel.TransactionTypeModel{
			Code: entity.GetTransactionType().GetCode(),
		},
		UpdatedAt: entity.GetUpdatedAt(),
	}

	statement, err := repository.client.Prepare("update categories set code = ?, name = ?, transaction_type_code = ?, updated_at = ? where id = ?")
	if err != nil {
		return status.InternalServerError, err
	}
	defer statement.Close()

	_, err = statement.Exec(model.Code, model.Name, model.TransactionType.Code, model.UpdatedAt, model.Id)
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		// Unique key violated
		return status.EntityWithSameKeyAlreadyExists, err
	}
	if err != nil {
		return status.InternalServerError, err
	}

	return status.Success, nil
}

func (repository CategoryRepository) FindById(id string) (entity.ICategory, error) {

	row, err := repository.client.Query("select categories.id, categories.tenant_id, categories.code, categories.name, transaction_types.code, transaction_types.name, categories.created_at, categories.updated_at from categories inner join transaction_types on categories.transaction_type_code = transaction_types.code where categories.id = ?", id)
	if err != nil {
		return entity.Category{}, err
	}
	defer row.Close()

	var model CategoryModel
	if row.Next() {
		if err := row.Scan(&model.Id, &model.TenantId, &model.Code, &model.Name, &model.TransactionType.Code, &model.TransactionType.Name, &model.CreatedAt, &model.UpdatedAt); err != nil {
			return entity.Category{}, err
		}

		transactionType, err := transactionTypeEntity.NewTransactionType(model.TransactionType.Code, model.TransactionType.Name)
		if err != nil {
			return entity.Category{}, err
		}
		category, err := entity.NewCategory(model.Id, model.TenantId, model.Code, model.Name, transactionType, model.CreatedAt, model.UpdatedAt)
		if err != nil {
			return entity.Category{}, err
		}
		return category, nil
	}
	return nil, nil
}

func (repository CategoryRepository) List(filterParameters []filter.FilterParameter, tenantId string) ([]entity.ICategory, error) {
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
		rows, err = repository.client.Query("select categories.id, categories.tenant_id, categories.code, categories.name, transaction_types.code, transaction_types.name, categories.created_at, categories.updated_at from categories inner join transaction_types on categories.transaction_type_code = transaction_types.code where categories.tenant_id = ?", tenantId)
	} else {
		if len(codeFilter) > 0 && len(nameFilter) == 0 {
			rows, err = repository.client.Query("select categories.id, categories.tenant_id, categories.code, categories.name, transaction_types.code, transaction_types.name, categories.created_at, categories.updated_at from categories inner join transaction_types on categories.transaction_type_code = transaction_types.code where categories.tenant_id = ? and code = ?", tenantId, codeFilter)
		}
		if len(codeFilter) == 0 && len(nameFilter) > 0 {
			rows, err = repository.client.Query("select categories.id, categories.tenant_id, categories.code, categories.name, transaction_types.code, transaction_types.name, categories.created_at, categories.updated_at from categories inner join transaction_types on categories.transaction_type_code = transaction_types.code where categories.tenant_id = ? and name = ?", tenantId, nameFilter)
		}
		if len(codeFilter) > 0 && len(nameFilter) > 0 {
			rows, err = repository.client.Query("select categories.id, categories.tenant_id, categories.code, categories.name, transaction_types.code, transaction_types.name, categories.created_at, categories.updated_at from categories inner join transaction_types on categories.transaction_type_code = transaction_types.code where categories.tenant_id = ? and code = ? and name = ?", tenantId, codeFilter, nameFilter)
		}
	}
	defer rows.Close()
	if err != nil {
		return []entity.ICategory{}, err
	}

	categories := []entity.ICategory{}
	for rows.Next() {
		var model CategoryModel

		if err := rows.Scan(&model.Id, &model.TenantId, &model.Code, &model.Name, &model.TransactionType.Code, &model.TransactionType.Name, &model.CreatedAt, &model.UpdatedAt); err != nil {
			return []entity.ICategory{}, err
		}

		transactionType, err := transactionTypeEntity.NewTransactionType(model.TransactionType.Code, model.TransactionType.Name)
		if err != nil {
			return []entity.ICategory{}, err
		}
		category, err := entity.NewCategory(model.Id, model.TenantId, model.Code, model.Name, transactionType, model.CreatedAt, model.UpdatedAt)
		if err != nil {
			return []entity.ICategory{}, err
		}

		categories = append(categories, category)
	}
	return categories, nil
}

func (repository CategoryRepository) Delete(id string) error {

	statement, err := repository.client.Prepare("delete from categories where id = ?")
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
