package main

import (
	"errors"
	"fmt"
	"nightblog/internal/cli"
	"nightblog/internal/config"
	"nightblog/internal/library"
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
	cfg, err := config.GetConfig()
	if err != nil {
		return nil, fmt.Errorf("Failed to load config: %v", err)
	}

	if err := library.VerifyLocalLibraryPath(cfg.LocalLibraryPath); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			if err := library.MakeLibraryDir(cfg.LocalLibraryPath); err != nil {
				return nil, fmt.Errorf("Failed to create library directory: %v", err)
			}
		} else {
			return nil, fmt.Errorf("Failed to verify library path: %v", err)
		}
	}

	return cfg, nil
}
