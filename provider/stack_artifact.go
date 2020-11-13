package provider

import (
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

type stackArtifact struct {
	ArtifactAccount string `mapstructure:"artifact_account"`
	ID              string `mapstructure:"id"`
	Name            string `mapstructure:"name"`
	Reference       string `mapstructure:"reference"`
	Type            string `mapstructure:"type"`
	Version         string `mapstructure:"version"`
}

func toClientStackArtifact(s *[]*stackArtifact) *client.StackArtifact {
	if s == nil || len(*s) == 0 {
		return nil
	}
	clientStackArtifact := client.StackArtifact(*(*s)[0])
	return &clientStackArtifact
}

func fromClientStackArtifact(cs *client.StackArtifact) *[]*stackArtifact {
	if cs == nil {
		return nil
	}
	s := stackArtifact(*cs)
	return &[]*stackArtifact{&s}
}
