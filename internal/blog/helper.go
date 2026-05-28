package blog

import "fmt"

// formatBlog generates the formatted blog post content including frontmatter and content.
func (b blog) formatBlog() string {
	frontmatter := fmt.Sprintf(
		"---\ntitle: %s\ndescription: %s\npublishedDate: %s\nupdatedDate: %s\n---\n",
		b.title,
		b.description,
		b.publishedDate,
		b.updatedDate,
	)
	blog := fmt.Sprintf("%s\n%s", frontmatter, b.content)
	return blog
}
