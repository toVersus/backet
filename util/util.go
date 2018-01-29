package util

import (
	"fmt"
	"os"
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
