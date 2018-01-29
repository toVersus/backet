package util

import (
	"fmt"
	"os/exec"
	"strings"
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

// IsGitInstalled returns true if git tool has been installed
func IsGitInstalled() bool {
	cmd := NewGitCommand("version")
	ver, err := exec.Command(cmd.name, cmd.args...).Output()
	if err != nil {
		return false
	}
	fmt.Printf("%s is installed\n", strings.TrimSpace(string(ver)))

	return true
}
