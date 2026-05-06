package cli

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
}

func createNewBlog(cfg *config.Config, title string) error {
	b := blog{
		title:         title,
		description:   "A short description",
		createdDate:   time.Now().Format(time.DateOnly),
		publishedDate: "",
		updatedDate:   time.Now().Format(time.DateOnly),
		content:       "# Blog content starts here",
	}

	title = strings.Replace(title, " ", "-", -1)
	file := filepath.Join(cfg.LocalLibraryPath, strings.ToLower(title)+".md")

	if _, err := os.Stat(file); err == nil {
		return fmt.Errorf("file %s already exists", file)
	}

	if err := os.WriteFile(file, []byte(b.formatBlog()), 0644); err != nil {
		return fmt.Errorf("failed to create file %s: %v", file, err)
	}

	return nil
}

func (b blog) formatBlog() string {
	frontmatter := fmt.Sprintf("---\ntitle: %s\ndescription: %s\npublishedDate: %s\nupdatedDate: %s\n---\n", b.title, b.description, b.publishedDate, b.updatedDate)
	blog := fmt.Sprintf("%s\n%s", frontmatter, b.content)
	return blog
}
