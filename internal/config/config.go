package config

import (
	"os"
	"path/filepath"

	"github.com/nightails/goconfig"
)

func configPath() string {
	cfgDir, _ := os.UserConfigDir()
	return filepath.Join(cfgDir, "nightblog")
}

func configFilePath() string {
	return filepath.Join(configPath(), "config.json")
}

type Config struct {
	RemoteBlogsURL string
	DatabasePath   string

	Editor string
}

func defaultConfig() *Config {
	cacheDir, _ := os.UserCacheDir()

	return &Config{
		RemoteBlogsURL: "",
		DatabasePath:   filepath.Join(cacheDir, "nightblog", "blogs.db"),
		Editor:         "nvim",
	}
}

func (cfg *Config) Load() error {
	return goconfig.Load(configFilePath(), cfg)
}

func (cfg *Config) Save() error {
	return goconfig.Save(configFilePath(), cfg)
}

func (cfg *Config) CreateDefault() error {
	cfg = defaultConfig()
	return goconfig.Create(configFilePath(), cfg)
}
