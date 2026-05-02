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
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return Config{}
	}

	return Config{
		LocalLibraryPath: filepath.Join(homeDir, "Documents", "Blogs"),
		RemoteLibraryURL: "",
	}
}

// GetConfig loads and parses the application's configuration from a JSON file, returning the resulting Config object.
func GetConfig() (*Config, error) {
	cfgFile, err := getConfigFile()
	if err != nil {
		return nil, err
	}
	defer cfgFile.Close()

	var cfg Config
	if err := json.NewDecoder(cfgFile).Decode(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

// Write saves the configuration to a JSON file in the user's home directory.
func (cfg *Config) Write() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	cfgFile := filepath.Join(homeDir, configPath, configFile)
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	if err := os.WriteFile(cfgFile, data, 0644); err != nil {
		return err
	}
	return nil
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
