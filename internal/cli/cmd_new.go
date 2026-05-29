package cli

import (
	"fmt"

	"nightblog/internal/blog"
	"nightblog/internal/config"

	"github.com/spf13/cobra"
)

func newCmd(cfg *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "new <title>",
		Short: "Create a new blog post",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := blog.New(cfg, args[0]); err != nil {
				return fmt.Errorf("failed to create blog post: %v", err)
			}

			fmt.Println("Blog post created successfully")
			return nil
		},
	}
}
