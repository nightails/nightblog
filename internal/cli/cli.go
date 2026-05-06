package cli

import (
	"fmt"
	"os"
)

func Run() {
	switch os.Args[1] {
	default:
		fmt.Printf("Unknown command: %s\n\n", os.Args[1])
	case "new":
		if len(os.Args) < 3 {
			fmt.Println("Missing title argument. Usage: new <title>")
			return
		}
		fmt.Println("Creating new blog...")
	}
}
