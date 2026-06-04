package storage

import (
	"database/sql"
)

func Init(dbPath string) (*sql.DB, *Queries, error) {
	// initialize SQLite and SQLC queries
	db, queries, err := initSQLite(dbPath)
	if err != nil {
		return nil, nil, err
	}

	// TODO: verify remote Database

	return db, queries, nil
}
