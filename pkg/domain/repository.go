package domain

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
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
