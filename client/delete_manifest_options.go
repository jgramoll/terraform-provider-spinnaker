package client

// DeleteManifestOptions options
type DeleteManifestOptions struct {
	Cascading bool `json:"cascading"`
}

// NewDeleteManifestOptions new options
func NewDeleteManifestOptions() *DeleteManifestOptions {
	return &DeleteManifestOptions{
		Cascading: true,
	}
}
