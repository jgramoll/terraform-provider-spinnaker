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

func (a *manifestInputArtifact) toClientInputArtifact() (*client.ManifestInputArtifact, error) {
	clientArtifact := &client.ManifestInputArtifact{}
	clientArtifact.ID = a.ID
	clientArtifact.Account = a.Account

	art, err := a.Artifact[0].toClientManifestArtifact()
	if err != nil {
		return nil, err
	}
	clientArtifact.Artifact = art

	return clientArtifact, nil
}
