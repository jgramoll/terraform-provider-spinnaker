package client

// NewExpectedArtifact new expected artifact
func NewExpectedArtifact() *ExpectedArtifact {
	return &ExpectedArtifact{}
}

// ExpectedArtifactDefaultArtifact default artifact
type ExpectedArtifactDefaultArtifact struct {
	CustomKind bool   `json:"customKind"`
	ID         string `json:"id"`
}

// ExpectedArtifactMatchArtifact match artifact
type ExpectedArtifactMatchArtifact struct {
	ID        string `json:"id"`
	Location  string `json:"location"`
	Name      string `json:"name"`
	Reference string `json:"reference"`
	Type      string `json:"type"`
}

// ExpectedArtifact artifacts expected from
type ExpectedArtifact struct {
	DefaultArtifact ExpectedArtifactDefaultArtifact `json:"defaultArtifact"`

	DisplayName string `json:"displayName"`
	ID          string `json:"id"`

	MatchArtifact ExpectedArtifactMatchArtifact `json:"matchArtifact"`

	UseDefaultArtifact bool `json:"useDefaultArtifact"`
	UsePriorArtifact   bool `json:"usePriorArtifact"`
}
