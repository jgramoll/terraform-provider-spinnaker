package provider

import (
	"github.com/google/uuid"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

type manifestExpectedArtifact struct {
	DefaultArtifact []manifestArtifact `mapstructure:"default_artifact"`

	DisplayName string `mapstructure:"display_name"`
	ID          string `mapstructure:"id"`

	MatchArtifact []manifestArtifact `mapstructure:"match_artifact"`

	UseDefaultArtifact bool `mapstructure:"use_default_artifact"`
	UsePriorArtifact   bool `mapstructure:"use_prior_artifact"`
}

func toClientExpectedArtifacts(artifacts *[]*manifestExpectedArtifact) (*[]*client.ManifestExpectedArtifact, error) {
	if artifacts == nil || len(*artifacts) == 0 {
		return nil, nil
	}
	newList := []*client.ManifestExpectedArtifact{}
	for _, a := range *artifacts {
		newExpectedArtifact := client.NewManifestExpectedArtifact()
		if a.DefaultArtifact != nil && len(a.DefaultArtifact) > 0 {
			clientArt, err := a.DefaultArtifact[0].toClientManifestArtifact()
			if err != nil {
				return nil, err
			}
			newExpectedArtifact.DefaultArtifact = *clientArt
		}
		newExpectedArtifact.DisplayName = a.DisplayName
		if a.ID != "" {
			newExpectedArtifact.ID = a.ID
		} else {
			id, err := uuid.NewRandom()
			if err != nil {
				return nil, err
			}
			newExpectedArtifact.ID = id.String()
		}
		if a.MatchArtifact != nil && len(a.MatchArtifact) > 0 {
			clientArt, err := a.MatchArtifact[0].toClientManifestArtifact()
			if err != nil {
				return nil, err
			}
			newExpectedArtifact.MatchArtifact = *clientArt
		}
		newExpectedArtifact.UseDefaultArtifact = a.UseDefaultArtifact
		newExpectedArtifact.UsePriorArtifact = a.UsePriorArtifact
		newList = append(newList, newExpectedArtifact)
	}
	return &newList, nil
}

func fromClientExpectedArtifacts(artifacts *[]*client.ManifestExpectedArtifact) *[]*manifestExpectedArtifact {
	if artifacts == nil {
		return nil
	}
	newList := []*manifestExpectedArtifact{}
	for _, a := range *artifacts {
		newExpectedArtifact := manifestExpectedArtifact{
			DefaultArtifact:    []manifestArtifact{manifestArtifact(a.DefaultArtifact)},
			DisplayName:        a.DisplayName,
			ID:                 a.ID,
			MatchArtifact:      []manifestArtifact{manifestArtifact(a.MatchArtifact)},
			UseDefaultArtifact: a.UseDefaultArtifact,
			UsePriorArtifact:   a.UsePriorArtifact,
		}
		newList = append(newList, &newExpectedArtifact)
	}
	return &newList
}
