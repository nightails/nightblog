package app

import (
	"database/sql"
	"fmt"

	"nightblog/internal/config"
	"nightblog/internal/storage"
)

type State struct {
	Config   *config.Config
	Database *sql.DB
	Queries  *storage.Queries
}

func Init() (*State, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	db, qr, err := storage.Init(cfg.DatabasePath)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize storage: %w", err)
	}

	return &State{
		Config:   cfg,
		Database: db,
		Queries:  qr,
	}, nil
}

func (s *State) Cleanup() {
	s.Database.Close()
}
