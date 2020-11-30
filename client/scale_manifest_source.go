package client

import "log"

// ScaleManifestSource source
type ScaleManifestSource int

const (
	// ScaleManifestSourceUnknown unknown
	ScaleManifestSourceUnknown ScaleManifestSource = iota
	// ScaleManifestSourceText text
	ScaleManifestSourceText
)

func (t ScaleManifestSource) String() string {
	return [...]string{"UNKNOWN", "text"}[t]
}

// ParseScaleManifestSource parse
func ParseScaleManifestSource(s string) (ScaleManifestSource, error) {
	switch s {
	default:
		log.Printf("[WARN] Unknown scale manifest source: %s\n", s)
		return ScaleManifestSourceUnknown, nil
	case "text":
		return ScaleManifestSourceText, nil
	}
}

// MarshalText marshal
func (t ScaleManifestSource) MarshalText() ([]byte, error) {
	return []byte(t.String()), nil
}

// UnmarshalText unmarshal
func (t *ScaleManifestSource) UnmarshalText(text []byte) error {
	source, err := ParseScaleManifestSource(string(text))
	if err != nil {
		return err
	}
	*t = source
	return nil
}
