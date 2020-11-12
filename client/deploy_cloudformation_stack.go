package client

// DeployCloudformationStackArtifact stack artifact
type DeployCloudformationStackArtifact struct {
	ArtifactAccount string `json:"artifactAccount"`
	ID              string `json:"id"`
	Name            string `json:"name"`
	Reference       string `json:"reference"`
	Type            string `json:"type"`
	Version         string `json:"version"`
}
