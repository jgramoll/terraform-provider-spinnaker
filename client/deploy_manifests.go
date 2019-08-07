package client

import (
	"github.com/ghodss/yaml"
)

type DeployManifests []string

func NewDeployManifests() *DeployManifests {
	return &DeployManifests{}
}

func (s DeployManifests) MarshalJSON() ([]byte, error) {
	jsonManifests := "["
	for _, manifest := range s {
		json, err := yaml.YAMLToJSON([]byte(manifest))
		if err != nil {
			return nil, err
		}
		jsonManifests += string(json)
	}
	return []byte(jsonManifests + "]"), nil
}

func ParseDeployManifests(manifestInterface []interface{}) (*DeployManifests, error) {
	manifests := NewDeployManifests()
	for _, manifest := range manifestInterface {
		b, err := yaml.Marshal(manifest)
		if err != nil {
			return nil, err
		}
		*manifests = append(*manifests, string(b))
	}
	return manifests, nil
}
