package client

import (
	"fmt"
	"strings"
)

// Manifests manifests
type Manifests []Manifest

// NewManifests new manifests
func NewManifests() *Manifests {
	return &Manifests{}
}

// MarshalJSON marshal
func (s Manifests) MarshalJSON() ([]byte, error) {
	var jsonManifests []string
	for _, manifest := range s {
		json, err := manifest.MarshalJSON()
		if err != nil {
			return nil, err
		}
		jsonManifests = append(jsonManifests, string(json))
	}
	jsonManifestsString := strings.Join(jsonManifests, ",")
	return []byte(fmt.Sprintf("[%s]", jsonManifestsString)), nil
}

// ParseManifests parse
func ParseManifests(manifestInterface []interface{}) (*Manifests, error) {
	manifests := NewManifests()
	for _, m := range manifestInterface {
		manifest, err := ParseManifest(m)
		if err != nil {
			return nil, err
		}
		*manifests = append(*manifests, Manifest(manifest))
	}
	return manifests, nil
}
