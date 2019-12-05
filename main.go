package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"os/exec"
	"time"

	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

// Secret represents the larger yaml object that is returned. We only need the data field.
type Secret struct {
	Value Data
}

// Data maps the data field of the secret's yaml to arbitrary key:value pairs of type string.
type Data struct {
	Data map[string]string `yaml:"data"`
}

var namespace string

// k8decodeCmd represents the k8decode command
var k8decodeCmd = &cobra.Command{
	Use:   "k8decode [SECRET-NAME] [OPTIONS]",
	Short: "Decodes the base64 encoded secrets from a Kubernetes secrets yaml output",
	Long:  `k8decode is a CLI tool for easily reading Kubernetes secrets without having to manually pipe individual secrets into a base64 decoder.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			log.Fatal("Error: please provide a valid secret name")
		}

		s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
		s.Color("green")
		s.Prefix = "[ DECODING SECRETS ] [ "
		s.Suffix = " ]"
		s.Start()

		baseArgs := args[0]

		out, err := exec.Command("kubectl", "get", "secret", "-n", namespace, baseArgs, "-o", "yaml").Output()
		if err != nil {
			log.Fatalf("Error getting secret " + err.Error())
		}

		m := Secret{}

		err = yaml.Unmarshal(out, &m.Value)
		if err != nil {
			log.Fatalf("Error parsing yaml %s", err)
		}
		fmt.Println()
		for k := range m.Value.Data {
			data, _ := base64.StdEncoding.DecodeString(m.Value.Data[k])
			fmt.Printf("%s : %v\n", k, string(data))
		}
		s.Stop()
	},
}

func main() {
	k8decodeCmd.Flags().StringVarP(&namespace, "namespace", "n", "default", "the namespace the secret lives in")
	k8decodeCmd.Execute()
}
