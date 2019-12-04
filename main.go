package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

type Secret struct {
	Value Data
}

type Data struct {
	Data map[string]string `yaml:"data"`
}

// k8decodeCmd represents the k8decode command
var k8decodeCmd = &cobra.Command{
	Use:   "k8decode",
	Short: "Decodes the base64 encoded secrets from a Kubernetes secrets yaml output",
	Long: `For example:
	Run: k8decode {name-of-kubernetes-secret}`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) <= 0 {
			log.Fatal("Error: please provide a valid secret name")
		}

		fmt.Println("Decoding base64")
		baseArgs := args[0]

		out, err := exec.Command("kubectl", "get", "secret", baseArgs, "-o", "yaml").Output()
		if err != nil {
			log.Fatalf("Error getting secret " + err.Error())
		}

		m := Secret{}

		err = yaml.Unmarshal(out, &m.Value)
		if err != nil {
			log.Fatalf("Error parsing yaml %s", err)
		}

		for k := range m.Value.Data {
			data, _ := base64.StdEncoding.DecodeString(m.Value.Data[k])
			fmt.Printf("%s : %v\n", k, string(data))

		}
	},
}

func main() {
	k8decodeCmd.GenBashCompletionFile("out.sh")
	k8decodeCmd.Execute()
}
