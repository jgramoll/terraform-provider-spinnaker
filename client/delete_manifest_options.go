package client

type DeleteManifestOptions struct {
	Cascading bool `json:"cascading"`
}

func NewDeleteManifestOptions() *DeleteManifestOptions {
	return &DeleteManifestOptions{
		Cascading: true,
	}
}
