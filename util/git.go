package util

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

// Command represents a command, its subcommands and arguments
type Command struct {
	name string
	args []string
	envs []string
}

func (c *Command) String() string {
	if len(c.args) == 0 {
		return c.name
	}
	return fmt.Sprintf("%s %s", c.name, strings.Join(c.args, " "))
}

// NewGitCommand returns a new Git command with given subcommand and arguments
func NewGitCommand(args ...string) *Command {
	return &Command{
		name: "git",
		args: args,
	}
}

// AddArgs adds new arguments to the command
func (c *Command) AddArgs(args ...string) *Command {
	c.args = append(c.args, args...)
	return c
}

// AddEnvs adds new arguments to the command
func (c *Command) AddEnvs(envs ...string) *Command {
	c.envs = append(c.envs, envs...)
	return c
}

//
func getGitVersion() string {
	cmd := NewGitCommand("version")
	v, _ := exec.Command(cmd.name, cmd.args...).Output()
	return strings.TrimSpace(string(v))
}

// IsGitInstalled returns true if git tool has been installed
func IsGitInstalled() bool {
	if ver := getGitVersion(); ver != "" {
		fmt.Printf("... %s is installed\n", ver)
		return true
	}
	return false
}

// RunInDir executes command in specified directory
func (c *Command) RunInDir(path string) ([]byte, error) {
	pwd, _ := filepath.Abs(".")

	if err := os.Chdir(path); err != nil {
		msg := fmt.Sprintf("could not move into %s\n", path)
		return nil, errors.Wrap(err, msg)
	}
	out, _ := exec.Command(c.name, c.args...).Output()

	if err := os.Chdir(pwd); err != nil {
		msg := fmt.Sprintf("could not move back to %s\n", pwd)
		return nil, errors.Wrap(err, msg)
	}

	return out, nil
}

// GitCloneWithMirrorOpt clones repository into dest directory with mirror option
func GitCloneWithMirrorOpt(src string, dest string) error {
	cmd := NewGitCommand("clone", "--mirror", src, dest)
	out, err := exec.Command(cmd.name, cmd.args...).Output()
	if err != nil {
		return errors.Wrap(err, "could not clone git repository\n")
	}
	fmt.Printf("  %s\n", out)
	return nil
}

// UpdateRepositoryInDir pulls changes from remote repository in path directory
// and goes back to current directory
func UpdateRepositoryInDir(path string) error {
	cmd := NewGitCommand("remote", "update")
	out, err := cmd.RunInDir(path)
	if err != nil {
		return errors.Wrap(err, "could not update repository\n")
	}
	fmt.Printf("  %s\n", out)
	return nil
}
