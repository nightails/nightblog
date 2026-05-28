package cli

import (
	"fmt"
	"nightblog/internal/blog"
	"nightblog/internal/config"

	"github.com/spf13/cobra"
)

func listCmd(cfg *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List blog posts",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			blogs, err := blog.List(cfg)
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
