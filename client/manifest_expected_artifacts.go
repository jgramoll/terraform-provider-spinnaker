package client

// ManifestExpectedArtifact artifacts expected from
type ManifestExpectedArtifact struct {
	DefaultArtifact    ManifestArtifact `json:"defaultArtifact"`
	DisplayName        string           `json:"displayName"`
	ID                 string           `json:"id"`
	MatchArtifact      ManifestArtifact `json:"matchArtifact"`
	UseDefaultArtifact bool             `json:"useDefaultArtifact"`
	UsePriorArtifact   bool             `json:"usePriorArtifact"`
}

// NewManifestExpectedArtifact new expected artifact
func NewManifestExpectedArtifact() *ManifestExpectedArtifact {
	return &ManifestExpectedArtifact{}
}
