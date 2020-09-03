package client

// ManifestInputArtifact bake manifest artifact
type ManifestInputArtifact struct {
	Account  string            `json:"account"`
	ID       string            `json:"id,omitempty"`
	Artifact *ManifestArtifact `json:"artifact,omitempty"`
}
