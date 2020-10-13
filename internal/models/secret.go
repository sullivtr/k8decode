package models

// Secret represents the larger yaml object that is returned. We only need the data field.
type Secret struct {
	Data map[string]string `yaml:"data"`
}
