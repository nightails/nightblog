package cli

import (
	"fmt"

	"nightblog/internal/app"
	"nightblog/internal/storage"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

func newCmd(s *app.State) *cobra.Command {
	return &cobra.Command{
		Use:   "new <title>",
		Short: "Create a new blog post",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			post, err := s.Queries.CreatePost(cmd.Context(), storage.CreatePostParams{
				ID:          uuid.New(),
				Title:       args[0],
				Description: "A default description",
				Content:     "Posts content go here",
				IsDraft:     true,
			})
			if err != nil {
				return err
			}

			fmt.Printf("Created post %v: %s\n", post.ID, post.Title)
			return nil
		},
	}
}
