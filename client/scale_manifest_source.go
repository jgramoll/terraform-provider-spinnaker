package client

import "log"

type ScaleManifestSource int

const (
	ScaleManifestSourceUnknown ScaleManifestSource = iota
	ScaleManifestSourceText
)

func (t ScaleManifestSource) String() string {
	return [...]string{"UNKNOWN", "text"}[t]
}

func ParseScaleManifestSource(s string) (ScaleManifestSource, error) {
	switch s {
	default:
		log.Println("[WARN] Unknown scale manifest source:", s)
		return ScaleManifestSourceUnknown, nil
	case "text":
		return ScaleManifestSourceText, nil
	}
}

func (t ScaleManifestSource) MarshalText() ([]byte, error) {
	return []byte(t.String()), nil
}

func (t *ScaleManifestSource) UnmarshalText(text []byte) error {
	source, err := ParseScaleManifestSource(string(text))
	if err != nil {
		return err
	}
	*t = source
	return nil
}
