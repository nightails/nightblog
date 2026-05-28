package cli

import (
	"fmt"
	"nightblog/internal/config"
	"os"

	"github.com/spf13/cobra"
)

func listCmd(cfg *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List blog posts",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			blogs, err := listBlogs(cfg)
			if err != nil {
				return fmt.Errorf("failed to list blogs: %w", err)
			}

			out := cmd.OutOrStdout()

			fmt.Fprintln(out, "Blogs:")
			for _, blog := range blogs {
				fmt.Fprintf(out, " - %s\n", blog)
			}

			return nil
		},
	}
}

func listBlogs(cfg *config.Config) ([]string, error) {
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
