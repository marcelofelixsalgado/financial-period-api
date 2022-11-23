package repository

import (
	"database/sql"
	"time"
)

type PeriodRepository struct {
	db *sql.DB
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
