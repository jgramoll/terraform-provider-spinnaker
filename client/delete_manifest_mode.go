package client

import (
	"errors"
	"fmt"
)

type DeleteManifestMode int

const (
	DeleteManifestModeUnknown DeleteManifestMode = iota
	DeleteManifestModeStatic
)

func (t DeleteManifestMode) String() string {
	return [...]string{"unknown", "static"}[t]
}

func ParseDeleteManifestMode(s string) (DeleteManifestMode, error) {
	switch s {
	default:
		return DeleteManifestModeUnknown, errors.New(fmt.Sprintf("Unknown Mode %s", s))
	case "static":
		return DeleteManifestModeStatic, nil
	}
}

func (m DeleteManifestMode) MarshalText() ([]byte, error) {
	return []byte(m.String()), nil
}

func (m *DeleteManifestMode) UnmarshalText(text []byte) error {
	mode, err := ParseDeleteManifestMode(string(text))
	if err != nil {
		return err
	}
	*m = mode
	return nil
}
