package cli

import (
	"nightblog/internal/app"

	"github.com/spf13/cobra"
)

func Execute(s *app.State) error {
	rootCmd := newRootCmd(s)
	return rootCmd.Execute()
}

func newRootCmd(s *app.State) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "nightblog",
		Short: "Manage blog posts from the terminal",
		Long:  "Nightblog is terminal application for creating, listing, and managing blog posts.",
		RunE: func(cmd *cobra.Command, args []string) error {
			// TODO: route to TUI when it is implemented
			// return tui.Run(cfg)
			return cmd.Help()
		},
	}

	return rootCmd
}
