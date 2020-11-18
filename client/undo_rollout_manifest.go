package client

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
)

// UndoRolloutManifests undo manifests
type UndoRolloutManifests []string

// NewUndoRolloutManifests new manifests
func NewUndoRolloutManifests() *UndoRolloutManifests {
	return &UndoRolloutManifests{}
}

// MarshalJSON marshal
func (s UndoRolloutManifests) MarshalJSON() ([]byte, error) {
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

// ParseUndoRolloutManifests parse
func ParseUndoRolloutManifests(manifestInterface []interface{}) (*UndoRolloutManifests, error) {
	manifests := NewUndoRolloutManifests()
	for _, manifest := range manifestInterface {
		b, err := yaml.Marshal(manifest)
		if err != nil {
			return nil, err
		}
		*manifests = append(*manifests, string(b))
	}
	return manifests, nil
}
