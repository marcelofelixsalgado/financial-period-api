package period

import (
	"database/sql"
	"errors"
	"time"

	"github.com/marcelofelixsalgado/financial-period-api/pkg/domain/period/entity"
	"github.com/marcelofelixsalgado/financial-period-api/pkg/infrastructure/repository/filter"

	"github.com/go-sql-driver/mysql"
	"github.com/marcelofelixsalgado/financial-commons/pkg/infrastructure/repository/status"
)

type PeriodRepository struct {
	client *sql.DB
}

type PeriodModel struct {
	id        string
	tenantId  string
	code      string
	name      string
	year      int
	startDate time.Time
	endDate   time.Time
	createdAt time.Time
	updatedAt time.Time
}

func NewPeriodRepository(client *sql.DB) IPeriodRepository {
	return &PeriodRepository{
		client: client,
	}
}

func (repository *PeriodRepository) Create(entity entity.IPeriod) (status.RepositoryInternalStatus, error) {
	var mysqlErr *mysql.MySQLError

	model := PeriodModel{
		id:        entity.GetId(),
		tenantId:  entity.GetTenantId(),
		code:      entity.GetCode(),
		name:      entity.GetName(),
		year:      entity.GetYear(),
		startDate: entity.GetStartDate(),
		endDate:   entity.GetEndDate(),
		createdAt: entity.GetCreatedAt(),
	}

	statement, err := repository.client.Prepare("insert into periods (id, tenant_id, code, name, year, start_date, end_date, created_at) values (?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return status.InternalServerError, err
	}
	defer statement.Close()

	_, err = statement.Exec(model.id, model.tenantId, model.code, model.name, model.year, model.startDate, model.endDate, model.createdAt)
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		// Unique key violated
		return status.EntityWithSameKeyAlreadyExists, err
	}
	if err != nil {
		return status.InternalServerError, err
	}

	return status.Success, nil
}

func (repository *PeriodRepository) FindById(id string) (entity.IPeriod, error) {

	row, err := repository.client.Query("select id, tenant_id, code, name, year, start_date, end_date, created_at, updated_at from periods where id = ?", id)
	if err != nil {
		return entity.Period{}, err
	}
	defer row.Close()

	var periodModel PeriodModel
	if row.Next() {
		if err := row.Scan(&periodModel.id, &periodModel.tenantId, &periodModel.code, &periodModel.name, &periodModel.year, &periodModel.startDate, &periodModel.endDate, &periodModel.createdAt, &periodModel.updatedAt); err != nil {
			return entity.Period{}, err
		}

		period, err := entity.NewPeriod(periodModel.id, periodModel.tenantId, periodModel.code, periodModel.name, periodModel.year, periodModel.startDate, periodModel.endDate, periodModel.createdAt, periodModel.updatedAt)
		if err != nil {
			return entity.Period{}, err
		}
		return period, nil
	}
	return nil, nil
}

func (repository *PeriodRepository) List(filterParameters []filter.FilterParameter, tenantId string) ([]entity.IPeriod, error) {

	codeFilter := ""
	nameFilter := ""
	dateFilter := ""
	for _, filterParameter := range filterParameters {
		switch filterParameter.Name {
		case "code":
			codeFilter = filterParameter.Value
		case "name":
			nameFilter = filterParameter.Value
		case "date":
			dateFilter = filterParameter.Value
		}
	}

	var rows *sql.Rows
	var err error
	if len(filterParameters) == 0 {
		rows, err = repository.client.Query("select id, tenant_id, code, name, year, start_date, end_date, created_at, updated_at from periods where tenant_id = ? order by start_date", tenantId)
	} else {
		if len(dateFilter) > 0 {
			rows, err = repository.client.Query("select id, tenant_id, code, name, year, start_date, end_date, created_at, updated_at from periods where tenant_id = ? and start_date <= ? and end_date >= ? order by start_date", tenantId, dateFilter, dateFilter)
		}
		if len(codeFilter) > 0 && len(nameFilter) == 0 {
			rows, err = repository.client.Query("select id, tenant_id, code, name, year, start_date, end_date, created_at, updated_at from periods where tenant_id = ? and code = ? order by start_date", tenantId, codeFilter)
		}
		if len(codeFilter) == 0 && len(nameFilter) > 0 {
			rows, err = repository.client.Query("select id, tenant_id, code, name, year, start_date, end_date, created_at, updated_at from periods where tenant_id = ? and name = ? order by start_date", tenantId, nameFilter)
		}
		if len(codeFilter) > 0 && len(nameFilter) > 0 {
			rows, err = repository.client.Query("select id, tenant_id, code, name, year, start_date, end_date, created_at, updated_at from periods where tenant_id = ? and code = ? and name = ? order by start_date", tenantId, codeFilter, nameFilter)
		}
	}

	if err != nil {
		return []entity.IPeriod{}, err
	}
	defer rows.Close()

	periods := []entity.IPeriod{}
	for rows.Next() {
		var periodModel PeriodModel

		if err := rows.Scan(&periodModel.id, &periodModel.tenantId, &periodModel.code, &periodModel.name, &periodModel.year, &periodModel.startDate, &periodModel.endDate, &periodModel.createdAt, &periodModel.updatedAt); err != nil {
			return []entity.IPeriod{}, err
		}

		period, err := entity.NewPeriod(periodModel.id, periodModel.tenantId, periodModel.code, periodModel.name, periodModel.year, periodModel.startDate, periodModel.endDate, periodModel.createdAt, periodModel.updatedAt)
		if err != nil {
			return []entity.IPeriod{}, err
		}

		periods = append(periods, period)
	}

	return periods, nil
}

func (repository *PeriodRepository) Update(entity entity.IPeriod) (status.RepositoryInternalStatus, error) {
	var mysqlErr *mysql.MySQLError

	model := PeriodModel{
		id:        entity.GetId(),
		tenantId:  entity.GetTenantId(),
		code:      entity.GetCode(),
		name:      entity.GetName(),
		year:      entity.GetYear(),
		startDate: entity.GetStartDate(),
		endDate:   entity.GetEndDate(),
		updatedAt: entity.GetUpdatedAt(),
	}

	statement, err := repository.client.Prepare("update periods set code = ?, name = ?, year = ?, start_date = ?, end_date = ?, updated_at = ? where id = ?")
	if err != nil {
		return status.InternalServerError, err
	}
	defer statement.Close()

	_, err = statement.Exec(model.code, model.name, model.year, model.startDate, model.endDate, model.updatedAt, model.id)
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		// Unique key violated
		return status.EntityWithSameKeyAlreadyExists, err
	}
	if err != nil {
		return status.InternalServerError, err
	}

	return status.Success, nil
}

func (repository *PeriodRepository) Delete(id string) error {

	statement, err := repository.client.Prepare("delete from periods where id = ?")
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

func (repository *PeriodRepository) FindOverlap(startDate time.Time, endDate time.Time, tenantId string) (status.RepositoryInternalStatus, error) {

	row, err := repository.client.Query("select id from periods where tenant_id = ? and (start_date between ? and ? or end_date between ? and ?)", tenantId, startDate, endDate, startDate, endDate)
	if err != nil {
		return status.InternalServerError, err
	}
	defer row.Close()

	if row.Next() {
		return status.OverlappingPeriodDates, nil
	}
	return status.Success, nil
}
