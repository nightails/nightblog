package cli

import (
	"fmt"
	"nightblog/internal/app"

	"github.com/spf13/cobra"
)

func listCmd(s *app.State) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List blog posts",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			posts, err := s.Queries.ListPosts(cmd.Context())
			if err != nil {
				return err
			}

			if len(posts) == 0 {
				fmt.Println("No posts found")
				return nil
			}

			for _, post := range posts {
				fmt.Printf("%d. %s\n", post.ID, post.Title)
			}
			return nil
		},
	}
}
