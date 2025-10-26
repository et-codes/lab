package driver

import (
	"database/sql"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
)

// DB holds the database connection pool
type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

const (
	maxOpenDBConn = 10
	maxIdleDBConn = 5
	maxDBLifetime = 5 * time.Minute
)

// ConnectSQL creates a db pool for Postgres
func ConnectSQL(dsn string) (*DB, error) {
	d, err := NewDataBase(dsn)
	if err != nil {
		return nil, err
	}

	d.SetMaxOpenConns(maxOpenDBConn)
	d.SetConnMaxIdleTime(maxIdleDBConn)
	d.SetConnMaxLifetime(maxDBLifetime)

	dbConn.SQL = d

	return dbConn, nil
}

// NewDataBase creates a new database
func NewDataBase(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
