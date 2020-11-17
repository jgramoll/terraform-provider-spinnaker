package client

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
)

// ScaleManifests scale
type ScaleManifests []string

// NewScaleManifests new scale
func NewScaleManifests() *ScaleManifests {
	return &ScaleManifests{}
}

// MarshalJSON marshal
func (s ScaleManifests) MarshalJSON() ([]byte, error) {
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

// ParseScaleManifests parse
func ParseScaleManifests(manifestInterface []interface{}) (*ScaleManifests, error) {
	manifests := NewScaleManifests()
	for _, manifest := range manifestInterface {
		b, err := yaml.Marshal(manifest)
		if err != nil {
			return nil, err
		}
		*manifests = append(*manifests, string(b))
	}
	return manifests, nil
}
