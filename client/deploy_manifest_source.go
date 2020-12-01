package client

import "log"

// DeployManifestSource source
type DeployManifestSource int

const (
	// DeployManifestSourceUnknown unknown
	DeployManifestSourceUnknown DeployManifestSource = iota
	// DeployManifestSourceText text
	DeployManifestSourceText
	// DeployManifestSourceArtifact artifact
	DeployManifestSourceArtifact
)

func (t DeployManifestSource) String() string {
	return [...]string{"UNKNOWN", "text", "artifact"}[t]
}

// ParseDeployManifestSource parse
func ParseDeployManifestSource(s string) (DeployManifestSource, error) {
	switch s {
	default:
		log.Printf("[WARN] Unknown deploy manifest source: %s\n", s)
		return DeployManifestSourceUnknown, nil
	case "text":
		return DeployManifestSourceText, nil
	case "artifact":
		return DeployManifestSourceArtifact, nil
	}
}

// MarshalText marshal
func (t DeployManifestSource) MarshalText() ([]byte, error) {
	return []byte(t.String()), nil
}

// UnmarshalText unmarshal
func (t *DeployManifestSource) UnmarshalText(text []byte) error {
	source, err := ParseDeployManifestSource(string(text))
	if err != nil {
		return err
	}
	*t = source
	return nil
}
