package secret

import (
	"encoding/base64"
	"fmt"
	"os/exec"
	"testing"

	"github.com/sullivtr/k8decode/internal/models"
	"gotest.tools/assert"
)

func TestGetSecretSuccess(t *testing.T) {
	s, _ := GetSecret(mockCmdRunnerSuccess{}, "testns", "testSecret")
	assert.Equal(t, s.Data["fakeSecretData"], "fakeSecretValue")
}

func TestGetSecretFail(t *testing.T) {
	_, err := GetSecret(mockCmdRunnerError{}, "testns", "testSecret")
	assert.ErrorContains(t, err, "Unable to execute command")
}

func TestGetSecretErrorUnmarshallingYaml(t *testing.T) {
	_, err := GetSecret(mockCmdRunnerUnmarshalError{}, "testns", "testSecret")
	assert.Error(t, err, "Error parsing yaml yaml: unmarshal errors:\n  line 2: cannot unmarshal !!str `notyaml` into models.Secret")
}

func TestPrintDecodedSecret(t *testing.T) {
	s := models.Secret{
		Data: map[string]string{},
	}
	s.Data["Key"] = base64.StdEncoding.EncodeToString([]byte("value"))

	PrintDecodedSecret(&s)
}

// Mock Cmd Runners
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

	out, err := exec.Command("notarealcommand").Output()
	return out, fmt.Errorf("Unable to execute command: %v", err)
}

type mockCmdRunnerUnmarshalError struct{}

func (r mockCmdRunnerUnmarshalError) Run(command string, args ...string) ([]byte, error) {
	secret := `
notyaml
`
	out := []byte(secret)
	return out, nil
}
