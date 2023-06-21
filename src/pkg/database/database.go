package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

// NewDatabase creates a new database connection based on the driver and dsn
// The dsn is the database connection string, path to the sqlite file or the postgres connection string, not required for memory
func NewDatabase(driver, dsn string) (*sql.DB, error) {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, err
	}
	return db, err
}

// Database is the database connection, it needs to be set before use
var Database *sql.DB
