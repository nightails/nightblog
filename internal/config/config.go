package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func configPath() string {
	cfgDir, _ := os.UserConfigDir()
	return filepath.Join(cfgDir, "nightblog")
}

func configFilePath() string {
	return filepath.Join(configPath(), "config.json")
}

type Config struct {
	LocalBlogsDir  string
	RemoteBlogsURL string
	DatabasePath   string

	Editor string
}

func defaultConfig() Config {
	homeDir, _ := os.UserHomeDir()
	cacheDir, _ := os.UserCacheDir()

	return Config{
		LocalBlogsDir:  filepath.Join(homeDir, "Documents", "Blogs"),
		RemoteBlogsURL: "",
		DatabasePath:   filepath.Join(cacheDir, "nightblog", "blogs.db"),
		Editor:         "nvim",
	}
}

// Load retrieves an existing configuration or generates a new default configuration if none exists.
func Load() (*Config, error) {
	file, err := openConfigFile()
	if err != nil {
		// TODO: replace with logger. May need to consolidate, check what being logged in the function first.
		return nil, fmt.Errorf("failed to open config file: %w", err)
	}
	defer file.Close()

	var cfg Config
	if err := json.NewDecoder(file).Decode(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

// Write saves the configuration to a JSON file in the user's home directory.
func Write(cfg *Config) error {
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	if err := os.WriteFile(configFilePath(), data, 0o755); err != nil {
		return err
	}
	return nil
}

// openConfigFile retrieves the configuration file for the application.
// NOTE: Close the file after use to prevent resource leaks.
func openConfigFile() (*os.File, error) {
	// create the config directory if not exist
	_, err := os.Stat(configPath())
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			if err := os.MkdirAll(configPath(), 0o755); err != nil {
				// TODO: replace with logger
				return nil, fmt.Errorf("failed to create config directory: %w", err)
			}
		} else {
			return nil, fmt.Errorf("failed to stat config path: %w", err)
		}
	}

	// open or create the config file if not exist
	file, err := os.Open(configFilePath())
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			file, err = os.Create(configFilePath())
			if err != nil {
				// TODO: replace with logger
				return nil, fmt.Errorf("failed to create config file: %w", err)
			}
			data, err := json.MarshalIndent(defaultConfig(), "", "  ")
			if err != nil {
				file.Close()
				// TODO: replace with logger
				return nil, fmt.Errorf("failed to marshal default config: %w", err)
			}
			// write default config to file
			_, err = file.Write(data)
			if err != nil {
				file.Close()
				// TODO: replace with logger
				return nil, fmt.Errorf("failed to write default config to file: %w", err)
			}
			// rest cursor to the beginging of the file
			if _, err := file.Seek(0, 0); err != nil {
				file.Close()
				return nil, fmt.Errorf("failed to seek the begining of the config file: %w", err)
			}
		} else {
			return nil, fmt.Errorf("failed to open config file: %w", err)
		}
	}

	return file, nil
}
