package provider

import (
	"github.com/google/uuid"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

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

func (a manifestArtifact) toClientManifestArtifact() (*client.ManifestArtifact, error) {
	clientArtifact := client.ManifestArtifact(a)

	if a.ID == "" {
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, err
		}
		clientArtifact.ID = id.String()
	}
	return &clientArtifact, nil
}
