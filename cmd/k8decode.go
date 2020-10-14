package cmd

import (
	"errors"
	"fmt"
	"log"

	"github.com/sullivtr/k8decode/internal/executor"
	"github.com/sullivtr/k8decode/internal/models"

	"github.com/spf13/cobra"
	"github.com/sullivtr/k8decode/internal/loader"
	"github.com/sullivtr/k8decode/internal/secret"
)

var namespace string
var cmdRunner executor.CommandRunner = executor.DefaultCommandRunner{}

var K8DecodeCmd = &cobra.Command{
	Use:   "k8decode",
	Short: "Decodes the base64 encoded secrets from a Kubernetes secrets yaml output",
	Long:  `k8decode is a CLI tool for easily reading Kubernetes secrets without having to manually pipe individual secrets into a base64 decoder.`,
	Args: func(cmd *cobra.Command, args []string) error {
		return validateArgs(args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := k8decode(args); err != nil {
			log.Fatalf("Error: %v", err)
		}
	},
}

func init() {
	K8DecodeCmd.Flags().StringVarP(&namespace, "namespace", "n", "default", "the namespace the secret lives in")
}

func validateArgs(args []string) error {
	if len(args) <= 0 || args[0] == "" {
		return errors.New("Error: please provide a valid secret name")
	}
	return nil
}

func k8decode(args []string) (*models.Secret, error) {
	l := loader.NewSpinner()
	l.Start()

	secretName := args[0]

	s, err := secret.GetSecret(cmdRunner, namespace, secretName)
	if err != nil {
		return nil, fmt.Errorf("Unable to get secret, %s, in namespace: %s. Error: %s", secretName, namespace, err.Error())
	}

	secret.PrintDecodedSecret(s)
	l.Stop()
	return s, nil
}
