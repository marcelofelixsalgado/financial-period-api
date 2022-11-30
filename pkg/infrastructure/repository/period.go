package repository

import (
	"database/sql"
	"marcelofelixsalgado/financial-period-api/pkg/domain/period/entity"
	"time"
)

type PeriodRepository struct {
	client *sql.DB
}

type PeriodModel struct {
	id        string
	code      string
	name      string
	year      int
	startDate time.Time
	endDate   time.Time
	createdAt time.Time
	updatedAt time.Time
}

func NewPeriodRepository(client *sql.DB) IRepository {
	return &PeriodRepository{
		client: client,
	}
}

func (repository *PeriodRepository) Create(entity entity.IPeriod) error {

	model := PeriodModel{
		id:        entity.GetId(),
		code:      entity.GetCode(),
		name:      entity.GetName(),
		year:      entity.GetYear(),
		startDate: entity.GetStartDate(),
		endDate:   entity.GetEndDate(),
		createdAt: entity.GetCreatedAt(),
	}

	statement, err := repository.client.Prepare("insert into periods (id, code, name, year, start_date, end_date, created_at) values (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(model.id, model.code, model.name, model.year, model.startDate, model.endDate, model.createdAt)
	if err != nil {
		return err
	}

	return nil
}

func (repository *PeriodRepository) FindById(id string) (entity.IPeriod, error) {

	row, err := repository.client.Query("select id, code, name, year, start_date, end_date, created_at, updated_at from periods where id = ?", id)
	if err != nil {
		return entity.Period{}, err
	}
	defer row.Close()

	var periodModel PeriodModel
	for row.Next() {
		if err := row.Scan(&periodModel.id, &periodModel.code, &periodModel.name, &periodModel.year, &periodModel.startDate, &periodModel.endDate, &periodModel.createdAt, &periodModel.updatedAt); err != nil {
			return entity.Period{}, err
		}
	}

	period, err := entity.NewPeriod(periodModel.id, periodModel.code, periodModel.name, periodModel.year, periodModel.startDate, periodModel.endDate, periodModel.createdAt, periodModel.updatedAt)
	if err != nil {
		return entity.Period{}, err
	}

	return period, nil
}

func (repository *PeriodRepository) FindAll(filterParameters []FilterParameter) ([]entity.IPeriod, error) {

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
	// fields := "id, code, name, year, start_date, end_date, created_at, updated_at"
	var rows *sql.Rows
	var err error
	if len(filterParameters) == 0 {
		rows, err = repository.client.Query("select id, code, name, year, start_date, end_date, created_at, updated_at from periods")
	} else {
		if len(codeFilter) > 0 && len(nameFilter) == 0 {
			rows, err = repository.client.Query("select id, code, name, year, start_date, end_date, created_at, updated_at from periods where code = ?", codeFilter)
		}
		if len(codeFilter) == 0 && len(nameFilter) > 0 {
			rows, err = repository.client.Query("select id, code, name, year, start_date, end_date, created_at, updated_at from periods where name = ?", nameFilter)
		}
		if len(codeFilter) > 0 && len(nameFilter) > 0 {
			rows, err = repository.client.Query("select id, code, name, year, start_date, end_date, created_at, updated_at from periods where code = ? and name = ?", codeFilter, nameFilter)
		}
	}

	if err != nil {
		return []entity.IPeriod{}, err
	}
	defer rows.Close()

	periods := []entity.IPeriod{}
	for rows.Next() {
		var periodModel PeriodModel

		if err := rows.Scan(&periodModel.id, &periodModel.code, &periodModel.name, &periodModel.year, &periodModel.startDate, &periodModel.endDate, &periodModel.createdAt, &periodModel.updatedAt); err != nil {
			return []entity.IPeriod{}, err
		}

		period, err := entity.NewPeriod(periodModel.id, periodModel.code, periodModel.name, periodModel.year, periodModel.startDate, periodModel.endDate, periodModel.createdAt, periodModel.updatedAt)
		if err != nil {
			return []entity.IPeriod{}, err
		}

		periods = append(periods, period)
	}

	return periods, nil
}

func (repository *PeriodRepository) Update(entity entity.IPeriod) error {

	model := PeriodModel{
		id:        entity.GetId(),
		code:      entity.GetCode(),
		name:      entity.GetName(),
		year:      entity.GetYear(),
		startDate: entity.GetStartDate(),
		endDate:   entity.GetEndDate(),
		updatedAt: entity.GetUpdatedAt(),
	}

	statement, err := repository.client.Prepare("update periods set code = ?, name = ?, year = ?, start_date = ?, end_date = ?, updated_at = ? where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(model.code, model.name, model.year, model.startDate, model.endDate, model.updatedAt, model.id)
	if err != nil {
		return err
	}

	return nil
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
