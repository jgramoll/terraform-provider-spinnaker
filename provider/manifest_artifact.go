package provider

type manifestArtifact struct {
	ArtifactAccount string            `mapstructure:"artifact_account"`
	CustomKind      bool              `mapstructure:"custom_kind"`
	ID              string            `mapstructure:"id"`
	Location        string            `mapstructure:"location"`
	Metadata        map[string]string `mapstructure:"metadata"`
	Name            string            `mapstructure:"name"`
	Reference       string            `mapstructure:"reference"`
	Type            string            `mapstructure:"type"`
	Version         string            `mapstructure:"version"`
}
