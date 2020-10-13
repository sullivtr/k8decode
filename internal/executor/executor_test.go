package executor

import (
	"testing"

	"gotest.tools/assert"
)

func TestRunCommand(t *testing.T) {
	runner := CmdRunner{}
	out, err := runner.Run("echo", "hello")
	assert.NilError(t, err)
	assert.Equal(t, string(out), "hello\n")
}
