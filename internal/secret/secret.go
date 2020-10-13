package secret

import (
	"encoding/base64"
	"fmt"

	"github.com/sullivtr/k8decode/internal/executor"
	"github.com/sullivtr/k8decode/internal/models"
	"gopkg.in/yaml.v2"
)

// GetSecret uses kubectl to fetch a secret in yaml format
func GetSecret(r executor.CommandRunner, namespace, secretName string) (*models.Secret, error) {
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

// PrintDecodedSecret will print the base64 decoded secrets line by line in the terminal
func PrintDecodedSecret(secret *models.Secret) {
	fmt.Println()
	for k := range secret.Data {
		data, _ := base64.StdEncoding.DecodeString(secret.Data[k])
		fmt.Printf("\033[32m%s : \033[37m%v\n", k, string(data))
	}
}
