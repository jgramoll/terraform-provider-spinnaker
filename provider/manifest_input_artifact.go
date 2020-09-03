package provider

import (
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

type manifestInputArtifact struct {
	Account  string             `mapstructure:"account"`
	ID       string             `mapstructure:"id"`
	Artifact []manifestArtifact `mapstructure:"artifact"`
}

func fromClientInputArtifact(ca *client.ManifestInputArtifact) *manifestInputArtifact {
	if ca == nil {
		return nil
	}

	return &manifestInputArtifact{
		Account:  ca.Account,
		ID:       ca.ID,
		Artifact: []manifestArtifact{manifestArtifact(*ca.Artifact)},
	}
}

func (a *manifestInputArtifact) toClientInputArtifact() *client.ManifestInputArtifact {
	art := client.ManifestArtifact(a.Artifact[0])
	return &client.ManifestInputArtifact{
		Account:  a.Account,
		ID:       a.ID,
		Artifact: &art,
	}
}
