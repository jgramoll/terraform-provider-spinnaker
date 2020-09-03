package provider

import (
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

func toClientExpectedArtifacts(artifacts *[]*manifestExpectedArtifact) *[]*client.ManifestExpectedArtifact {
	if artifacts == nil || len(*artifacts) == 0 {
		return nil
	}
	newList := []*client.ManifestExpectedArtifact{}
	for _, a := range *artifacts {
		newExpectedArtifact := client.NewManifestExpectedArtifact()
		if a.DefaultArtifact != nil && len(a.DefaultArtifact) > 0 {
			newExpectedArtifact.DefaultArtifact = client.ManifestArtifact(a.DefaultArtifact[0])
		}
		newExpectedArtifact.DisplayName = a.DisplayName
		newExpectedArtifact.ID = a.ID
		if a.MatchArtifact != nil && len(a.MatchArtifact) > 0 {
			newExpectedArtifact.MatchArtifact = client.ManifestArtifact(a.MatchArtifact[0])
		}
		newExpectedArtifact.UseDefaultArtifact = a.UseDefaultArtifact
		newExpectedArtifact.UsePriorArtifact = a.UsePriorArtifact
		newList = append(newList, newExpectedArtifact)
	}
	return &newList
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
