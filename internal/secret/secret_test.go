package secret

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/sullivtr/k8decode/internal/executor"
	"github.com/sullivtr/k8decode/internal/models"
	"github.com/sullivtr/k8decode/internal/reader"
)

type SecretTestSuite struct {
	suite.Suite
}

func TestSecretTestSuite(t *testing.T) {
	suite.Run(t, new(SecretTestSuite))
}

func (suite SecretTestSuite) TestGetSecretSuccess() {
	s, _ := GetSecret(mockCmdRunnerSuccess{}, "testns", "testSecret")
	suite.Equal(s.Data["fakeSecretData"], "fakeSecretValue")
}

func (suite SecretTestSuite) TestGetSecretFail() {
	_, err := GetSecret(executor.DefaultCommandRunner{}, "testns", "")
	suite.Contains(err.Error(), "Error getting secret")
	suite.Error(err)
}

func (suite SecretTestSuite) TestGetSecretErrorUnmarshallingYaml() {
	_, err := GetSecret(mockCmdRunnerUnmarshalError{}, "testns", "testSecret")
	suite.Contains(err.Error(), "Error parsing yaml yaml: unmarshal errors:\n  line 2: cannot unmarshal !!str `notyaml` into models.Secret")
	suite.Error(err)
}

func (suite SecretTestSuite) TestPrintDecodedSecret() {
	s := models.Secret{
		Data: map[string]string{},
	}
	s.Data["Key"] = base64.StdEncoding.EncodeToString([]byte("value"))

	res := reader.ReadStdOut(func() {
		PrintDecodedSecret(&s)
	})
	suite.Contains(res, "Key")
	suite.Contains(res, "value")
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

type mockCmdRunnerUnmarshalError struct{}

func (r mockCmdRunnerUnmarshalError) Run(command string, args ...string) ([]byte, error) {
	secret := `
notyaml
`
	out := []byte(secret)
	return out, nil
}
