package cli

import (
	"nightblog/internal/config"

	"github.com/spf13/cobra"
)

func Execute(cfg *config.Config) error {
	rootCmd := newRootCmd(cfg)
	return rootCmd.Execute()
}

func newRootCmd(cfg *config.Config) *cobra.Command {
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

	rootCmd.AddCommand(newCmd(cfg))
	rootCmd.AddCommand(listCmd(cfg))

	return rootCmd
}
