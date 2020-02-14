package client

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
)

type EnableServerGroups []string

func NewEnableServerGroups() *EnableServerGroups {
	return &EnableServerGroups{}
}

func (s EnableServerGroups) MarshalJSON() ([]byte, error) {
	var jsonManifests []string
	for _, manifest := range s {
		json, err := yaml.YAMLToJSON([]byte(manifest))
		if err != nil {
			return nil, err
		}
		jsonManifests = append(jsonManifests, string(json))
	}
	jsonManifestsString := strings.Join(jsonManifests, ",")
	return []byte(fmt.Sprintf("[%s]", jsonManifestsString)), nil
}

func ParseEnableServerGroups(manifestInterface []interface{}) (*EnableServerGroups, error) {
	manifests := NewEnableServerGroups()
	for _, manifest := range manifestInterface {
		b, err := yaml.Marshal(manifest)
		if err != nil {
			return nil, err
		}
		*manifests = append(*manifests, string(b))
	}
	return manifests, nil
}
