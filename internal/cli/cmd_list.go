package cli

import (
	"nightblog/internal/app"

	"github.com/spf13/cobra"
)

func listCmd(s *app.State) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List blog posts",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
}
