package util

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
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

// IsGitInstalled returns true if git tool has been installed
func IsGitInstalled() bool {
	ver, err := exec.Command("git", "version").Output()
	if err != nil {
		return false
	}
	fmt.Printf("%s is installed\n", strings.TrimSpace(string(ver)))

	return true
}
