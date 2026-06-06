package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	_ "modernc.org/sqlite"

	"nightblog/internal/app"
	"nightblog/internal/cli"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	state, err := app.Init()
	if err != nil {
		fmt.Printf("Initialization failed: %v\n", err)
		return
	}
	defer state.Cleanup()

	if err := cli.Execute(ctx, state); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
