package client

// ManifestArtifact bake manifest artifact data
type ManifestArtifact struct {
	ArtifactAccount string            `json:"artifactAccount,omitempty"`
	CustomKind      bool              `json:"customKind"`
	ID              string            `json:"id"`
	Location        string            `json:"location,omitempty"`
	Metadata        map[string]string `json:"metadata,omitempty"`
	Name            string            `json:"name,omitempty"`
	Reference       string            `json:"reference,omitempty"`
	Type            string            `json:"type,omitempty"`
	Version         string            `json:"version,omitempty"`
}
