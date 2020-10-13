package executor

import "os/exec"

// CommandRunner represents the base for a command executor
type CommandRunner interface {
	Run(string, ...string) ([]byte, error)
}

// CmdRunner is the implementation of Runner, which wraps exec.Command
type CmdRunner struct{}

// Run executes a command with args
func (r CmdRunner) Run(command string, args ...string) ([]byte, error) {

	out, err := exec.Command(command, args...).Output()

	return out, err

}
