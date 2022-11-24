package database

import (
	"database/sql"
	"log"
	"marcelofelixsalgado/financial-period-api/configs"
)

var ConnectionPool *sql.DB

func Connect() {
	db, err := sql.Open("mysql", configs.DatabaseConnectionString)
	if err != nil {
		log.Fatalf("Error trying to connect to database: %v", err)
	}

	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(3)
	db.SetMaxOpenConns(3)

	// Checks if connection is open
	if err = db.Ping(); err != nil {
		db.Close()
		log.Fatalf("Error trying to check the database connection: %v", err)
	}

	ConnectionPool = db
}
