package cmd

import (
	"encoding/base64"
	"fmt"
	"os/exec"
	"testing"

	"github.com/sullivtr/k8decode/internal/models"

	"gotest.tools/assert"
)

func TestGetSecretSuccess(t *testing.T) {
	s, _ := getSecret(MockCmdRunnerSuccess{}, "testns", "testSecret")
	assert.Equal(t, s.Data["fakeSecretData"], "fakeSecretValue")
}

func TestGetSecretFail(t *testing.T) {
	_, err := getSecret(MockCmdRunnerError{}, "testns", "testSecret")
	assert.ErrorContains(t, err, "Unable to execute command")
}

func TestGetSecretErrorUnmarshallingYaml(t *testing.T) {
	_, err := getSecret(MockCmdRunnerUnmarshalError{}, "testns", "testSecret")
	assert.Error(t, err, "Error parsing yaml yaml: unmarshal errors:\n  line 2: cannot unmarshal !!str `notyaml` into models.Secret")
}

func TestPrintDecodedSecret(t *testing.T) {
	s := models.Secret{
		Data: map[string]string{},
	}
	s.Data["Key"] = base64.StdEncoding.EncodeToString([]byte("value"))

	printDecodedSecret(&s)
}

type MockCmdRunnerSuccess struct{}

func (r MockCmdRunnerSuccess) Run(command string, args ...string) ([]byte, error) {
	secret := `
data:
  fakeSecretData: fakeSecretValue
`
	out := []byte(secret)
	return out, nil
}

type MockCmdRunnerError struct{}

func (r MockCmdRunnerError) Run(command string, args ...string) ([]byte, error) {

	out, err := exec.Command("notarealcommand").Output()
	return out, fmt.Errorf("Unable to execute command: %v", err)
}

type MockCmdRunnerUnmarshalError struct{}

func (r MockCmdRunnerUnmarshalError) Run(command string, args ...string) ([]byte, error) {
	secret := `
notyaml
`
	out := []byte(secret)
	return out, nil
}
