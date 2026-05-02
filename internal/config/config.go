package config

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

const configPath = ".config/blogs_manager"
const configFile = "config.json"

type Config struct {
	LocalLibraryPath string
	RemoteLibraryURL string
}

func defaultConfig() Config {
	return Config{
		LocalLibraryPath: "$HOME/Documents/blogs/",
		RemoteLibraryURL: "",
	}
}

// GetConfigFile retrieves the configuration file for the application.
// NOTE: Close the file after use to prevent resource leaks.
func getConfigFile() (*os.File, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	cfgDir := filepath.Join(homeDir, configPath)
	cfgFilePath := filepath.Join(cfgDir, configFile)

	file, err := os.Open(cfgFilePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			if err := os.MkdirAll(cfgDir, 0755); err != nil {
				return nil, err
			}

			data, err := json.MarshalIndent(defaultConfig(), "", "  ")
			if err != nil {
				return nil, err
			}

			if err := os.WriteFile(cfgFilePath, data, 0644); err != nil {
				return nil, err
			}

			file, err = os.Open(cfgFilePath)
			if err != nil {
				return nil, err
			}

			return file, nil
		}
		return nil, err
	}

	return file, nil
}
