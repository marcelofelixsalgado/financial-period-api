package repository

import (
	"database/sql"
	"marcelofelixsalgado/financial-period-api/pkg/domain/period/entity"

	_ "github.com/go-sql-driver/mysql"
)

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

func (model PeriodModel) Update() error {
	return nil
}

func (model PeriodModel) Find(id string) (entity.Period, error) {
	return entity.Period{}, nil
}

func (model PeriodModel) FindAll() ([]entity.Period, error) {
	return nil, nil
}

func (model PeriodModel) Delete(id string) error {
	return nil
}

func NewRepository() IRepository {
	return PeriodModel{}
}
