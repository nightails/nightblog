package app

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"nightblog/internal/config"
	"nightblog/internal/storage"
)

type State struct {
	Config   *config.Config
	Database *sql.DB
	Queries  *storage.Queries
}

func Init() (*State, error) {
	cfg := &config.Config{}
	if err := cfg.Load(); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			if err = cfg.CreateDefault(); err != nil {
				return nil, err
			}
		}
		return nil, err
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
