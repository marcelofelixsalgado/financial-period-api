package subcategory

import (
	"database/sql"
	"errors"
	"time"

	categoryEntity "github.com/marcelofelixsalgado/financial-period-api/pkg/domain/category/entity"
	"github.com/marcelofelixsalgado/financial-period-api/pkg/domain/subcategory/entity"
	transactionTypeEntity "github.com/marcelofelixsalgado/financial-period-api/pkg/domain/transactiontype/entity"
	"github.com/marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/category"
	"github.com/marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/filter"

	"github.com/marcelofelixsalgado/financial-commons/pkg/infrastructure/repository/status"

	"github.com/go-sql-driver/mysql"
)

type SubCategoryRepository struct {
	client *sql.DB
}

type SubCategoryModel struct {
	id        string
	tenantId  string
	code      string
	name      string
	category  category.CategoryModel
	createdAt time.Time
	updatedAt time.Time
}

func NewSubCategoryRepository(client *sql.DB) ISubCategoryRepository {
	return &SubCategoryRepository{
		client: client,
	}
}

func (repository SubCategoryRepository) Create(entity entity.ISubCategory) (status.RepositoryInternalStatus, error) {
	var mysqlErr *mysql.MySQLError

	model := SubCategoryModel{
		id:       entity.GetId(),
		tenantId: entity.GetTenantId(),
		code:     entity.GetCode(),
		name:     entity.GetName(),
		category: category.CategoryModel{
			Id: entity.GetCategory().GetId(),
		},
		createdAt: entity.GetCreatedAt(),
		updatedAt: entity.GetUpdatedAt(),
	}

	statement, err := repository.client.Prepare("insert into subcategories (id, tenant_id, code, name, category_id, created_at) values (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return status.InternalServerError, err
	}
	defer statement.Close()

	_, err = statement.Exec(model.id, model.tenantId, model.code, model.name, model.category.Id, model.createdAt)
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		// Unique key violated
		return status.EntityWithSameKeyAlreadyExists, err
	}
	if err != nil {
		return status.InternalServerError, err
	}

	return status.Success, nil
}

func (repository SubCategoryRepository) Update(entity entity.ISubCategory) (status.RepositoryInternalStatus, error) {
	var mysqlErr *mysql.MySQLError

	model := SubCategoryModel{
		id:       entity.GetId(),
		tenantId: entity.GetTenantId(),
		code:     entity.GetCode(),
		name:     entity.GetName(),
		category: category.CategoryModel{
			Id: entity.GetCategory().GetId(),
		},
		createdAt: entity.GetCreatedAt(),
		updatedAt: entity.GetUpdatedAt(),
	}

	statement, err := repository.client.Prepare("update subcategories set code = ?, name = ?, category_id = ?, updated_at = ? where id = ?")
	if err != nil {
		return status.InternalServerError, err
	}
	defer statement.Close()

	_, err = statement.Exec(model.code, model.name, model.category.Id, model.updatedAt, model.id)
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		// Unique key violated
		return status.EntityWithSameKeyAlreadyExists, err
	}
	if err != nil {
		return status.InternalServerError, err
	}

	return status.Success, nil
}

func (repository SubCategoryRepository) FindById(id string) (entity.ISubCategory, error) {

	row, err := repository.client.Query("select subcategories.id, subcategories.tenant_id, subcategories.code, subcategories.name, categories.id, categories.code, categories.name, transaction_types.code, transaction_types.name, subcategories.created_at, subcategories.updated_at from subcategories inner join categories on subcategories.category_id = categories.id inner join transaction_types on categories.transaction_type_code = transaction_types.code where subcategories.id = ?", id)
	if err != nil {
		return entity.SubCategory{}, nil
	}
	defer row.Close()

	if row.Next() {
		var model SubCategoryModel

		if err := row.Scan(&model.id, &model.tenantId, &model.code, &model.name, &model.category.Id, &model.category.Code, &model.category.Name, &model.category.TransactionType.Code, &model.category.TransactionType.Name, &model.createdAt, &model.updatedAt); err != nil {
			return entity.SubCategory{}, err
		}

		transactionType, err := transactionTypeEntity.NewTransactionType(model.category.TransactionType.Code, model.category.TransactionType.Name)
		if err != nil {
			return entity.SubCategory{}, err
		}
		category, err := categoryEntity.NewCategory(model.category.Id, model.tenantId, model.category.Code, model.category.Name, transactionType, time.Time{}, time.Time{})
		if err != nil {
			return entity.SubCategory{}, err
		}
		subCategory, err := entity.NewSubCategory(model.id, model.tenantId, model.code, model.name, category, model.createdAt, model.updatedAt)
		if err != nil {
			return entity.SubCategory{}, err
		}
		return subCategory, nil
	}
	return nil, nil
}

func (repository SubCategoryRepository) List(filterParameters []filter.FilterParameter, tenantId string) ([]entity.ISubCategory, error) {

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
		rows, err = repository.client.Query("select subcategories.id, subcategories.tenant_id, subcategories.code, subcategories.name, categories.id, categories.code, categories.name, transaction_types.code, transaction_types.name, subcategories.created_at, subcategories.updated_at from subcategories inner join categories on subcategories.category_id = categories.id inner join transaction_types on categories.transaction_type_code = transaction_types.code where subcategories.tenant_id = ?", tenantId)
	} else {
		if len(codeFilter) > 0 && len(nameFilter) == 0 {
			rows, err = repository.client.Query("select subcategories.id, subcategories.tenant_id, subcategories.code, subcategories.name, categories.id, categories.code, categories.name, transaction_types.code, transaction_types.name, subcategories.created_at, subcategories.updated_at from subcategories inner join categories on subcategories.category_id = categories.id inner join transaction_types on categories.transaction_type_code = transaction_types.code where subcategories.tenant_id = ? and subcategories.code = ?", tenantId, codeFilter)
		}
		if len(codeFilter) == 0 && len(nameFilter) > 0 {
			rows, err = repository.client.Query("select subcategories.id, subcategories.tenant_id, subcategories.code, subcategories.name, categories.id, categories.code, categories.name, transaction_types.code, transaction_types.name, subcategories.created_at, subcategories.updated_at from subcategories inner join categories on subcategories.category_id = categories.id inner join transaction_types on categories.transaction_type_code = transaction_types.code where subcategories.tenant_id = ? and subcategories.name = ?", tenantId, nameFilter)
		}
		if len(codeFilter) > 0 && len(nameFilter) > 0 {
			rows, err = repository.client.Query("select subcategories.id, subcategories.tenant_id, subcategories.code, subcategories.name, categories.id, categories.code, categories.name, transaction_types.code, transaction_types.name, subcategories.created_at, subcategories.updated_at from subcategories inner join categories on subcategories.category_id = categories.id inner join transaction_types on categories.transaction_type_code = transaction_types.code where subcategories.tenant_id = ? and subcategories.code = ? and subcategories.name = ?", tenantId, codeFilter, nameFilter)
		}
	}
	defer rows.Close()
	if err != nil {
		return []entity.ISubCategory{}, err
	}

	subCategories := []entity.ISubCategory{}
	for rows.Next() {
		var model SubCategoryModel

		if err := rows.Scan(&model.id, &model.tenantId, &model.code, &model.name, &model.category.Id, &model.category.Code, &model.category.Name, &model.category.TransactionType.Code, &model.category.TransactionType.Name, &model.createdAt, &model.updatedAt); err != nil {
			return []entity.ISubCategory{}, err
		}

		transactionType, err := transactionTypeEntity.NewTransactionType(model.category.TransactionType.Code, model.category.TransactionType.Name)
		if err != nil {
			return []entity.ISubCategory{}, err
		}
		category, err := categoryEntity.NewCategory(model.category.Id, model.tenantId, model.category.Code, model.category.Name, transactionType, time.Time{}, time.Time{})
		if err != nil {
			return []entity.ISubCategory{}, err
		}
		subCategory, err := entity.NewSubCategory(model.id, model.tenantId, model.code, model.name, category, model.createdAt, model.updatedAt)
		if err != nil {
			return []entity.ISubCategory{}, err
		}
		subCategories = append(subCategories, subCategory)
	}

	return subCategories, nil
}

func (repository SubCategoryRepository) Delete(id string) error {

	statement, err := repository.client.Prepare("delete from subcategories where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(id); err != nil {
		return err
	}

	return nil
}
