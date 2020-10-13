package runner

import "os/exec"

// Runner represents the base for a command executor
type Runner interface {
	Run(string, ...string) ([]byte, error)
}

// RunCommand is the implementation of Runner, which wraps exec.Command
type RunCommand struct{}

// Run executes a command with args
func (r RunCommand) Run(command string, args ...string) ([]byte, error) {

	out, err := exec.Command(command, args...).Output()

	return out, err

}
