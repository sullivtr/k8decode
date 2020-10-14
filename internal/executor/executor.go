package executor

import "os/exec"

// CommandRunner represents the base for a command executor
type CommandRunner interface {
	Run(string, ...string) ([]byte, error)
}

// DefaultCommandRunner is the implementation of Runner, which wraps exec.Command
type DefaultCommandRunner struct{}

var _ CommandRunner = DefaultCommandRunner{}

// Run executes a command with args
func (r DefaultCommandRunner) Run(command string, args ...string) ([]byte, error) {

	out, err := exec.Command(command, args...).Output()

	return out, err

}
