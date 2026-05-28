package blog

import (
	"fmt"
	"nightblog/internal/config"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type blog struct {
	title         string
	description   string
	createdDate   string
	publishedDate string
	updatedDate   string
	content       string
	draft         bool
}

// New creates a new blog post with the given title and configuration.
func New(cfg *config.Config, title string) error {
	b := blog{
		title:         title,
		description:   "A short description",
		createdDate:   time.Now().Format(time.DateOnly),
		publishedDate: "",
		updatedDate:   time.Now().Format(time.DateOnly),
		content:       "# Blog content starts here",
		draft:         true,
	}

	title = strings.Replace(title, " ", "-", -1)
	file := filepath.Join(cfg.LocalBlogsDir, strings.ToLower(title)+".md")

	if _, err := os.Stat(file); err == nil {
		return fmt.Errorf("file %s already exists", file)
	}

	if err := os.WriteFile(file, []byte(b.formatBlog()), 0644); err != nil {
		return fmt.Errorf("failed to create file %s: %v", file, err)
	}

	return nil
}

// List retrieves a list of blog post filenames from the directory specified in the provided configuration.
func List(cfg *config.Config) ([]string, error) {
	entries, err := os.ReadDir(cfg.LocalBlogsDir)
	if err != nil {
		return nil, err
	}

	blogs := make([]string, 0, len(entries))
	for _, entry := range entries {
		blogs = append(blogs, entry.Name())
	}

	return blogs, nil
}
