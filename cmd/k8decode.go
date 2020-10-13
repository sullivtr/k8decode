package cmd

import (
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"
	"github.com/sullivtr/k8decode/internal/models"
	"github.com/sullivtr/k8decode/internal/runner"
	"gopkg.in/yaml.v2"
)

var namespace string
var cmdRunner *runner.Runner

var K8DecodeCmd = &cobra.Command{
	Use:   "k8decode",
	Short: "Decodes the base64 encoded secrets from a Kubernetes secrets yaml output",
	Long:  `k8decode is a CLI tool for easily reading Kubernetes secrets without having to manually pipe individual secrets into a base64 decoder.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) <= 0 {
			return errors.New("Error: please provide a valid secret name")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		s := initSpinner()
		s.Start()

		secretName := args[0]

		secret, err := getSecret(runner.RunCommand{}, namespace, secretName)
		if err != nil {
			log.Fatalf("Unable to get secret, %s, in namespace: %s", secretName, namespace)
		}

		// fmt.Println()
		printDecodedSecret(secret)
		s.Stop()
	},
}

func init() {
	K8DecodeCmd.Flags().StringVarP(&namespace, "namespace", "n", "default", "the namespace the secret lives in")
}

func initSpinner() *spinner.Spinner {
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Color("green")
	s.Prefix = "[ DECODING SECRETS ] [ "
	s.Suffix = " ]"

	return s
}

func getSecret(r runner.Runner, namespace, secretName string) (*models.Secret, error) {
	out, err := r.Run("kubectl", "get", "secret", "-n", namespace, secretName, "-o", "yaml")
	if err != nil {
		return nil, fmt.Errorf("Error getting secret " + err.Error())
	}

	m := models.Secret{}

	err = yaml.Unmarshal(out, &m)
	if err != nil {
		return nil, fmt.Errorf("Error parsing yaml %s", err)
	}
	return &m, nil
}

func printDecodedSecret(secret *models.Secret) {
	fmt.Println()
	for k := range secret.Data {
		data, _ := base64.StdEncoding.DecodeString(secret.Data[k])
		fmt.Printf("\033[32m%s : \033[37m%v\n", k, string(data))
	}
}
