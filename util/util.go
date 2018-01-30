package util

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

// ErrorExit outputs message on stdout and exit with status code 1
func ErrorExit(msg string) {
	fmt.Printf(msg)
	os.Exit(1)
}

// IsExist returns true if specified file/directory path exists
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// CreateDir creates a new directory if not exists
func CreateDir(path string) error {
	fmt.Printf("Creating new directory: %s\n", path)
	if err := os.MkdirAll(path, 0666); err != nil {
		return errors.Wrap(err, "could not create a new directory")
	}
	return nil
}

// GetParentDirName returns a parent direcotory name from input path
func GetParentDirName(path string) string {
	var tmp string
	for {
		if path != "." {
			tmp = path
			path = filepath.Dir(path)
		} else {
			break
		}
	}
	return tmp
}

// DeleteDir deletes a directory with its child directory and files
func DeleteDir(path string) error {
	fmt.Printf("Deleting directory: %s\n", path)
	if err := os.RemoveAll(path); err != nil {
		return errors.Wrap(err, "could not delete directory")
	}
	return nil
}
