package storage

import (
	"database/sql"

	"nightblog/internal/config"
)

type Storage struct {
	Database *sql.DB
	Queries  *Queries
}

func Init(cfg *config.Config) (*Storage, error) {
	// initialize SQLite and SQLC queries
	db, queries, err := initSQLite(cfg.DatabasePath)
	if err != nil {
		return nil, err
	}

	// TODO: verify remote Database

	return &Storage{db, queries}, nil
}

func (stg *Storage) CleanUp() {
	stg.Database.Close()
}
