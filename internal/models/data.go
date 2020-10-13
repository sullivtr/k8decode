package models

// Data maps the data field of the secret's yaml to arbitrary key:value pairs of type string.
type Data struct {
	Data map[string]string `yaml:"data"`
}
