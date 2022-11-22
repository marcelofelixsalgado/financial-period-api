package repository

import (
	"database/sql"
	"marcelofelixsalgado/financial-period-api/pkg/domain/period/entity"

	_ "github.com/go-sql-driver/mysql"
)

func NewRepository() IRepository {
	return PeriodModel{}
}

func connect() (*sql.DB, error) {
	connectionUrl := "root:root@tcp(financial-db:3306)/financial_db?charset=utf8&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", connectionUrl)
	if err != nil {
		return nil, err
	}

	// Checks if connection is open
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func (model PeriodModel) Create(entity entity.IPeriod) error {

	model = PeriodModel{
		id:        entity.GetId(),
		code:      entity.GetCode(),
		name:      entity.GetName(),
		year:      entity.GetYear(),
		startDate: entity.GetStartDate(),
		endDate:   entity.GetEndDate(),
		createdAt: entity.GetCreatedAt(),
		updatedAt: entity.GetUpdatedAt(),
	}

	db, err := connect()
	if err != nil {
		return err
	}
	defer db.Close()

	statement, err := db.Prepare("insert into periods (id, code, name, year, start_date, end_date, created_at) values (?, ?, ?, ?, ?, ?, ?)")
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

func (model PeriodModel) FindById(id string) (entity.IPeriod, error) {
	db, err := connect()
	if err != nil {
		return entity.Period{}, err
	}
	defer db.Close()

	row, err := db.Query("select id, code, name, year, start_date, end_date, created_at, updated_at from periods where id = ?", id)
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

func (model PeriodModel) FindAll() ([]entity.IPeriod, error) {

	db, err := connect()
	if err != nil {
		return []entity.IPeriod{}, err
	}
	defer db.Close()

	rows, err := db.Query("select id, code, name, year, start_date, end_date, created_at, updated_at from periods")
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

func (model PeriodModel) Delete(id string) error {
	return nil
}

func (model PeriodModel) Update(entity entity.IPeriod) error {

	model = PeriodModel{
		id:        entity.GetId(),
		code:      entity.GetCode(),
		name:      entity.GetName(),
		year:      entity.GetYear(),
		startDate: entity.GetStartDate(),
		endDate:   entity.GetEndDate(),
		updatedAt: entity.GetUpdatedAt(),
	}

	db, err := connect()
	if err != nil {
		return err
	}
	defer db.Close()

	statement, err := db.Prepare("update periods set code = ?, name = ?, year = ?, start_date = ?, end_date = ?, updated_at = ? where id = ?")
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
