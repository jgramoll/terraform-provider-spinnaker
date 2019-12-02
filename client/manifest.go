package client

import (
	"fmt"

	"github.com/ghodss/yaml"
)

type Manifest string

func (s Manifest) MarshalJSON() ([]byte, error) {
	jsonManifest, err := yaml.YAMLToJSON([]byte(s))
	if err != nil {
		return nil, err
	}
	return []byte(fmt.Sprintf("%s", jsonManifest)), nil
}

func ParseManifest(manifestInterface interface{}) (Manifest, error) {
	manifest, err := yaml.Marshal(manifestInterface)
	if err != nil {
		return "", err
	}
	return Manifest(manifest), nil
}
