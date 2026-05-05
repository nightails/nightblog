package library

import (
	"errors"
	"fmt"
	"os"
)

func VerifyLocalLibraryPath(path string) error {
	stat, err := os.Stat(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return err
		}
		return err
	}
	if !stat.IsDir() {
		return fmt.Errorf("path %s is not a directory", path)
	}
	return nil
}

func MakeLibraryDir(path string) error {
	return os.MkdirAll(path, 0755)
}
