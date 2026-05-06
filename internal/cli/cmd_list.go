package cli

import (
	"nightblog/internal/config"
	"os"
)

func listBlogs(cfg *config.Config) ([]string, error) {
	entries, err := os.ReadDir(cfg.LocalLibraryPath)
	if err != nil {
		return nil, err
	}

	blogs := make([]string, 0, len(entries))
	for _, entry := range entries {
		blogs = append(blogs, entry.Name())
	}

	return blogs, nil
}
