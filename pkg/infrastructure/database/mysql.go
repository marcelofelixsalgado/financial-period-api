package database

import (
	"database/sql"
	"marcelofelixsalgado/financial-period-api/configs"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", configs.DatabaseConnectionString)
	if err != nil {
		return nil, err
	}

	// Checks if connection is open
	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
