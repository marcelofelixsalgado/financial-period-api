package balance

import (
	"database/sql"
	"errors"
	"time"

	"github.com/marcelofelixsalgado/financial-period-api/pkg/domain/balance/entity"

	"github.com/marcelofelixsalgado/financial-commons/pkg/infrastructure/repository/status"

	"github.com/go-sql-driver/mysql"
)

type BalanceRepository struct {
	client *sql.DB
}

type BalanceModel struct {
	id           string
	tenantId     string
	periodId     string
	categoryId   string
	actualAmount float32
	limitAmount  float32
	createdAt    time.Time
	updatedAt    time.Time
}

func NewBalanceRepository(client *sql.DB) IBalanceRepository {
	return &BalanceRepository{
		client: client,
	}
}

func (repository *BalanceRepository) Create(entity entity.IBalance) (status.RepositoryInternalStatus, error) {
	var mysqlErr *mysql.MySQLError

	model := BalanceModel{
		id:           entity.GetId(),
		tenantId:     entity.GetTenantId(),
		periodId:     entity.GetPeriodId(),
		categoryId:   entity.GetCategoryId(),
		actualAmount: entity.GetActualAmount(),
		limitAmount:  entity.GetLimitAmount(),
		createdAt:    entity.GetCreatedAt(),
		updatedAt:    entity.GetUpdatedAt(),
	}

	statement, err := repository.client.Prepare("insert into balance (id, tenant_id, period_id, category_id, actual_amount, limit_amount, created_at) values (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return status.InternalServerError, err
	}
	defer statement.Close()

	_, err = statement.Exec(model.id, model.tenantId, model.periodId, model.categoryId, model.actualAmount, model.limitAmount, model.createdAt)
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		// Unique key violated
		return status.EntityWithSameKeyAlreadyExists, err
	}
	if err != nil {
		return status.InternalServerError, err
	}

	return status.Success, nil
}

func (repository *BalanceRepository) Update(entity entity.IBalance) error {

	model := BalanceModel{
		id:           entity.GetId(),
		tenantId:     entity.GetTenantId(),
		periodId:     entity.GetPeriodId(),
		categoryId:   entity.GetCategoryId(),
		actualAmount: entity.GetActualAmount(),
		limitAmount:  entity.GetLimitAmount(),
		createdAt:    entity.GetCreatedAt(),
		updatedAt:    entity.GetUpdatedAt(),
	}

	statement, err := repository.client.Prepare("update balance set period_id = ?, category_id = ?, actual_amount = ?, limit_amount = ?, updated_at = ? where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(model.periodId, model.categoryId, model.actualAmount, model.limitAmount, model.updatedAt, model.id)
	if err != nil {
		return err
	}

	return nil
}

func (repository *BalanceRepository) FindById(id string) (entity.IBalance, error) {

	row, err := repository.client.Query("select id, tenant_id, period_id, category_id, actual_amount, limit_amount, created_at, updated_at from balance where id = ?", id)
	if err != nil {
		return entity.Balance{}, err
	}
	defer row.Close()

	var balanceModel BalanceModel
	if row.Next() {
		if err := row.Scan(&balanceModel.id, &balanceModel.tenantId, &balanceModel.periodId, &balanceModel.categoryId, &balanceModel.actualAmount, &balanceModel.limitAmount, &balanceModel.createdAt, &balanceModel.updatedAt); err != nil {
			return entity.Balance{}, err
		}

		balance, err := entity.NewBalance(balanceModel.id, balanceModel.tenantId, balanceModel.periodId, balanceModel.categoryId, balanceModel.actualAmount, balanceModel.limitAmount, balanceModel.createdAt, balanceModel.updatedAt)
		if err != nil {
			return entity.Balance{}, err
		}
		return balance, nil
	}

	return nil, nil
}

func (repository *BalanceRepository) List(tenantId string, periodId string) ([]entity.IBalance, error) {

	var rows *sql.Rows
	var err error
	if periodId != "" {
		rows, err = repository.client.Query("select id, tenant_id, period_id, category_id, actual_amount, limit_amount from balance where tenant_id = ? and period_id = ?", tenantId, periodId)
	} else {
		rows, err = repository.client.Query("select id, tenant_id, period_id, category_id, actual_amount, limit_amount from balance where tenant_id = ?", tenantId)
	}

	if err != nil {
		return []entity.IBalance{}, err
	}
	defer rows.Close()

	balances := []entity.IBalance{}
	for rows.Next() {
		var balanceModel BalanceModel

		if err := rows.Scan(&balanceModel.id, &balanceModel.tenantId, &balanceModel.periodId, &balanceModel.categoryId, &balanceModel.actualAmount, &balanceModel.limitAmount); err != nil {
			return []entity.IBalance{}, err
		}

		balance, err := entity.NewBalance(balanceModel.id, balanceModel.tenantId, balanceModel.periodId, balanceModel.categoryId, balanceModel.actualAmount, balanceModel.limitAmount, balanceModel.createdAt, balanceModel.updatedAt)
		if err != nil {
			return []entity.IBalance{}, err
		}

		balances = append(balances, balance)
	}

	return balances, nil
}

func (repository *BalanceRepository) Delete(id string) error {

	statement, err := repository.client.Prepare("delete from balance where id = ?")
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
