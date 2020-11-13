package client

import (
	"fmt"
)

// DeployCloudformationSource source type
type DeployCloudformationSource int

const (
	// DeployCloudformationSourceUnknown default unknown
	DeployCloudformationSourceUnknown DeployCloudformationSource = iota
	// DeployCloudformationSourceText text
	DeployCloudformationSourceText
	// DeployCloudformationSourceArtifact artifact
	DeployCloudformationSourceArtifact
)

func (s DeployCloudformationSource) String() string {
	return [...]string{"unknown", "text", "artifact"}[s]
}

// ParseDeployCloudformationSource parse source
func ParseDeployCloudformationSource(s string) (DeployCloudformationSource, error) {
	switch s {
	default:
		return DeployCloudformationSourceUnknown, fmt.Errorf("Unknown source %s", s)
	case "text":
		return DeployCloudformationSourceText, nil
	case "artifact":
		return DeployCloudformationSourceArtifact, nil
	}
}

// MarshalText source to text
func (s DeployCloudformationSource) MarshalText() ([]byte, error) {
	return []byte(s.String()), nil
}

// UnmarshalText source from text
func (s *DeployCloudformationSource) UnmarshalText(text []byte) error {
	parsedSource, err := ParseDeployCloudformationSource(string(text))
	if err != nil {
		return err
	}
	*s = parsedSource
	return nil
}
