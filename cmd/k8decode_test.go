package cmd

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestValidateArgsSuccess(t *testing.T) {
	args := make([]string, 1)
	args[0] = "arg1"
	err := validateArgs(args)

	assert.NilError(t, err)
}

func TestValidateArgsError(t *testing.T) {
	args := make([]string, 1)
	args[0] = ""
	err := validateArgs(args)

	assert.Error(t, err, "Error: please provide a valid secret name")
}

func TestK8DecodeSuccess(t *testing.T) {
	args := make([]string, 2)
	args[0] = "testSecret"
	args[1] = "testns"
	cmdRunner = mockCmdRunnerSuccess{}

	s, _ := k8decode(args)
	assert.Equal(t, s.Data["fakeSecretData"], "fakeSecretValue")
}

func TestK8DecodeError(t *testing.T) {
	args := make([]string, 2)
	args[0] = "testSecret"
	cmdRunner = mockCmdRunnerError{}

	_, err := k8decode(args)
	assert.ErrorContains(t, err, "Unable to get secret, testSecret, in namespace: default")
}

type mockCmdRunnerSuccess struct{}

func (r mockCmdRunnerSuccess) Run(command string, args ...string) ([]byte, error) {
	secret := `
data:
  fakeSecretData: fakeSecretValue
`
	out := []byte(secret)
	return out, nil
}

type mockCmdRunnerError struct{}

func (r mockCmdRunnerError) Run(command string, args ...string) ([]byte, error) {
	return nil, fmt.Errorf("Test Error")
}
