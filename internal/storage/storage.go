package storage

import (
	"errors"
	"fmt"
	"nightblog/internal/config"
	"os"
)

func VerifyOrCrate(cfg *config.Config) error {
	if err := verifyLocalStorage(cfg.LocalBlogsDir); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			if err := makeStoragePath(cfg.LocalBlogsDir); err != nil {
				return fmt.Errorf("failed to create storage directory: %v", err)
			}
		} else {
			return fmt.Errorf("failed to verify storage path: %v", err)
		}
	}
	// TODO: verify remote Database
	return nil
}

func verifyLocalStorage(path string) error {
	stat, err := os.Stat(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return err
		}
		return err
	}
	if !stat.IsDir() {
		return fmt.Errorf("path %s is not a directory", path)
	}
	return nil
}

func makeStoragePath(path string) error {
	return os.MkdirAll(path, 0755)
}
