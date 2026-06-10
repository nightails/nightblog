package cli

import (
	"fmt"
	"os"
	"os/exec"

	"nightblog/internal/app"
	"nightblog/internal/storage"

	"github.com/spf13/cobra"
)

func editCmd(s *app.State) *cobra.Command {
	return &cobra.Command{
		Use:   "edit <title>",
		Short: "Edit the content of an exiting blog post",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			post, err := s.Queries.GetPostByTitle(cmd.Context(), args[0])
			if err != nil {
				return err
			}

			f, err := os.CreateTemp("", "tmp_post.*.md")
			if err != nil {
				return err
			}
			defer os.Remove(f.Name())

			f.Write([]byte(post.Content))
			// close temp file to prevent `file in use` conflict
			f.Close()

			if err := openEditor(s.Config.Editor, f.Name()); err != nil {
				return fmt.Errorf("editor failed to open: %w", err)
			}

			content, err := os.ReadFile(f.Name())
			if err != nil {
				return err
			}

			s.Queries.UpdatePostContent(
				cmd.Context(),
				storage.UpdatePostContentParams{
					ID:      post.ID,
					Content: string(content),
				},
			)

			return nil
		},
	}
}

func openEditor(editorPath string, filePath string) error {
	cmd := exec.Command(editorPath, filePath)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
