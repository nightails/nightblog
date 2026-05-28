package main

import (
	"fmt"
	"nightblog/internal/cli"
	"nightblog/internal/config"
	"nightblog/internal/storage"
	"os"
)

func main() {
	cfg, err := initApp()
	if err != nil {
		fmt.Printf("Initialization failed: %v\n", err)
		return
	}

	if len(os.Args) < 2 {
		// TODO: route to TUI
		fmt.Println("Missing argument, TUI is not implemented yet")
		return
	}

	cli.Run(cfg)
}

func initApp() (*config.Config, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %v", err)
	}

	if err := storage.Init(cfg); err != nil {
		return nil, fmt.Errorf("failed to initialize storage: %v", err)
	}

	return cfg, nil
}
