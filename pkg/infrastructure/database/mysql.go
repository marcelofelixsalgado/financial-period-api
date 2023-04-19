package database

import (
	"database/sql"
	"fmt"

	"github.com/marcelofelixsalgado/financial-period-api/commons/logger"
	"github.com/marcelofelixsalgado/financial-period-api/settings"

	_ "github.com/go-sql-driver/mysql"
)

func NewConnection() *sql.DB {

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		settings.Config.DatabaseConnectionUser,
		settings.Config.DatabaseConnectionPassword,
		settings.Config.DatabaseConnectionServerAddress,
		settings.Config.DatabaseConnectionServerPort,
		settings.Config.DatabaseName)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		logger.GetLogger().Fatalf("Error trying to connect to database: %v", err)
	}

	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(3)
	db.SetMaxOpenConns(3)

	// Checks if connection is open
	if err = db.Ping(); err != nil {
		db.Close()
		logger.GetLogger().Fatalf("Error trying to check the database connection: %v", err)
	}

	return db
}
