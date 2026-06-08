package cli

import (
	"nightblog/internal/app"

	"github.com/spf13/cobra"
)

func removeCmd(s *app.State) *cobra.Command {
	return &cobra.Command{
		Use:   "remove <title>",
		Short: "Remove a blog post for given title",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			post, err := s.Queries.GetPostByTitle(cmd.Context(), args[0])
			if err != nil {
				return err
			}

			if err = s.Queries.RemovePost(cmd.Context(), post.ID); err != nil {
				return err
			}
			return nil
		},
	}
}
