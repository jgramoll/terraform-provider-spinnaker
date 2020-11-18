package client

import (
	"fmt"

	"github.com/ghodss/yaml"
)

// Manifest manifest
type Manifest string

// MarshalJSON marshal
func (s Manifest) MarshalJSON() ([]byte, error) {
	jsonManifest, err := yaml.YAMLToJSON([]byte(s))
	if err != nil {
		return nil, err
	}
	return []byte(fmt.Sprintf("%s", jsonManifest)), nil
}

// ParseManifest parse
func ParseManifest(manifestInterface interface{}) (Manifest, error) {
	manifest, err := yaml.Marshal(manifestInterface)
	if err != nil {
		return "", err
	}
	return Manifest(manifest), nil
}
