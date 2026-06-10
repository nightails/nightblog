package cli

import (
	"os"
	"os/exec"

	"nightblog/internal/app"

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
			defer f.Close()

			f.Write([]byte(post.Content))
			openEditor(s.Config.Editor, f.Name())

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
