package cli

import (
	"fmt"
	"nightblog/internal/config"
	"os"
)

func Run(cfg *config.Config) {
	switch os.Args[1] {
	default:
		fmt.Printf("Unknown command: %s\n\n", os.Args[1])

	case "new":
		if len(os.Args) < 3 {
			fmt.Println("Missing title argument. Usage: new <title>")
			return
		}
		if err := createNewBlog(cfg, os.Args[2]); err != nil {
			fmt.Printf("Failed to create new blog: %v\n", err)
			return
		}
		fmt.Println("Blog created successfully!")

	case "list":
		blogs, err := listBlogs(cfg)
		if err != nil {
			fmt.Printf("Failed to list blogs: %v\n", err)
			return
		}
		fmt.Println("Blogs:")
		for _, blog := range blogs {
			fmt.Printf(" - %s\n", blog)
		}
	}
}
