package main

import (
	"blogs_manager/internal/config"
	"blogs_manager/internal/library"
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	fmt.Println("Config loaded successfully")
	fmt.Printf("Config: %+v\n", cfg)

	if err := library.VerifyLocalLibraryPath(cfg.LocalLibraryPath); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			if err := library.MakeLibraryDir(cfg.LocalLibraryPath); err != nil {
				log.Fatalf("Failed to create library directory: %v", err)
			}
		} else {
			log.Fatalf("Failed to verify library path: %v", err)
		}
	}
	fmt.Println("Library path verified or created successfully")
}
