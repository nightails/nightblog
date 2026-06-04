package main

import (
	"fmt"
	"os"

	"nightblog/internal/cli"
	"nightblog/internal/config"
	"nightblog/internal/storage"
)

func main() {
	cfg, stg, err := initApp()
	if err != nil {
		fmt.Printf("Initialization failed: %v\n", err)
		return
	}
	defer stg.CleanUp()

	if err := cli.Execute(cfg, stg); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initApp() (*config.Config, *storage.Storage, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to load config: %v", err)
	}

	stg, err := storage.Init(cfg)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to initialize storage: %v", err)
	}

	return cfg, stg, nil
}
