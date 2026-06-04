package main

import (
	"fmt"
	"os"

	"nightblog/internal/app"
	"nightblog/internal/cli"
)

func main() {
	state, err := app.Init()
	if err != nil {
		fmt.Printf("Initialization failed: %v\n", err)
		return
	}
	defer state.Cleanup()

	if err := cli.Execute(state); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
