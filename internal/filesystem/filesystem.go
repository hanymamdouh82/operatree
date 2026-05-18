package filesystem

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

// if found return with error, any other case return nil
func CheckDirExists(pth string) error {

	_, err := os.Stat(pth)

	// exists, return with error
	if err == nil {
		return fmt.Errorf("dir exists")
	}

	// doesn't exist -> return with nil error
	if errors.Is(err, fs.ErrNotExist) {
		return nil
	}

	// default case, cannot identify error -> assume exists and return with error
	return err
}

func CreateDir(pth string) error {

	if err := os.Mkdir(pth, 0775); err != nil {
		return err
	}

	return nil
}
