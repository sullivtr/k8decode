// +build test

package cmd

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/sullivtr/k8decode/internal/executor"
	"github.com/sullivtr/k8decode/internal/models"
	"github.com/sullivtr/k8decode/mocks"
	"gopkg.in/yaml.v2"
)

const (
	validYaml = `
data:
  fakeSecretData: fakeSecretValue
`
)

type K8decodeTestSuite struct {
	suite.Suite
}

func TestK8decodeSuite(t *testing.T) {
	suite.Run(t, new(K8decodeTestSuite))
}

func (suite K8decodeTestSuite) TestValidateArgsSuccess() {
	args := make([]string, 1)
	args[0] = "arg1"
	err := validateArgs(args)
	suite.Nil(err)
}

func (suite K8decodeTestSuite) TestValidateArgsError() {
	args := make([]string, 1)
	args[0] = ""
	err := validateArgs(args)
	suite.Error(err)
}

func (suite K8decodeTestSuite) TestK8DecodeSuccess() {
	cmdr := new(mocks.CommandRunner)
	namespace = "testns"
	runnerArgs := []interface{}{"kubectl", "get", "secret", "-n", "testns", "testsecret", "-o", "yaml"}
	cmdr.On("Run", runnerArgs...).Return([]byte(validYaml), nil)
	cmdRunner = cmdr
	cliArgs := []string{"testsecret"}
	s, err := k8decode(cliArgs)
	suite.Nil(err)

	d := models.Secret{}
	_ = yaml.Unmarshal([]byte(validYaml), &d)
	suite.Equal(s.Data["fakeSecretData"], d.Data["fakeSecretData"])

	cmdr.AssertExpectations(suite.T())
}

func (suite K8decodeTestSuite) TestK8DecodeError() {
	cmdr := new(mocks.CommandRunner)
	namespace = "testns"
	runnerArgs := []interface{}{"kubectl", "get", "secret", "-n", "testns", "testsecret", "-o", "yaml"}
	cmdr.On("Run", runnerArgs...).Return(nil, fmt.Errorf("TestError"))
	cmdRunner = cmdr
	cliArgs := []string{"testsecret"}
	_, err := k8decode(cliArgs)
	suite.NotNil(err)
	cmdr.AssertExpectations(suite.T())
}

func (suite K8decodeTestSuite) TestK8DecodeNoSecretName() {
	namespace = "testns"
	cliArgs := []string{""}
	cmdRunner = executor.DefaultCommandRunner{}
	_, err := k8decode(cliArgs)
	suite.Contains(err.Error(), "Unable to get secret, , in namespace: testns. Error: Error getting secret")
	suite.NotNil(err)
}
