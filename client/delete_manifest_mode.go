package client

import (
	"fmt"
)

// DeleteManifestMode delete manfiest mode
type DeleteManifestMode int

const (
	// DeleteManifestModeUnknown unknown
	DeleteManifestModeUnknown DeleteManifestMode = iota
	// DeleteManifestModeStatic static
	DeleteManifestModeStatic
)

func (m DeleteManifestMode) String() string {
	return [...]string{"unknown", "static"}[m]
}

// ParseDeleteManifestMode mode
func ParseDeleteManifestMode(s string) (DeleteManifestMode, error) {
	switch s {
	default:
		return DeleteManifestModeUnknown, fmt.Errorf("Unknown Mode %s", s)
	case "static":
		return DeleteManifestModeStatic, nil
	}
}

// MarshalText marshal text
func (m DeleteManifestMode) MarshalText() ([]byte, error) {
	return []byte(m.String()), nil
}

// UnmarshalText unmarshal text
func (m *DeleteManifestMode) UnmarshalText(text []byte) error {
	mode, err := ParseDeleteManifestMode(string(text))
	if err != nil {
		return err
	}
	*m = mode
	return nil
}
