package main

import (
	"blogs_manager/internal/config"
	"fmt"
	"log"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	fmt.Println("Config loaded successfully")
	fmt.Printf("Config: %+v\n", cfg)
}
