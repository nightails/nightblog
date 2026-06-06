package cli

import (
	"nightblog/internal/config"

	"github.com/spf13/cobra"
)

func newCmd(cfg *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "new <title>",
		Short: "Create a new blog post",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
}
