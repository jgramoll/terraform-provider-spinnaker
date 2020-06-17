package provider

import (
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

type expectedArtifactDefaultArtifact struct {
	CustomKind bool   `mapstructure:"custom_kind"`
	ID         string `mapstructure:"id"`
}

type expectedArtifactMatchArtifact struct {
	ID        string `mapstructure:"id"`
	Location  string `mapstructure:"location"`
	Name      string `mapstructure:"name"`
	Reference string `mapstructure:"reference"`
	Type      string `mapstructure:"type"`
}

type expectedArtifact struct {
	DefaultArtifact []expectedArtifactDefaultArtifact `mapstructure:"default_artifact"`

	DisplayName string `mapstructure:"display_name"`
	ID          string `mapstructure:"id"`

	MatchArtifact []expectedArtifactMatchArtifact `mapstructure:"match_artifact"`

	UseDefaultArtifact bool `mapstructure:"use_default_artifact"`
	UsePriorArtifact   bool `mapstructure:"use_prior_artifact"`
}

func toClientExpectedArtifacts(artifacts *[]*expectedArtifact) *[]*client.ExpectedArtifact {
	if artifacts == nil || len(*artifacts) == 0 {
		return nil
	}
	newList := []*client.ExpectedArtifact{}
	for _, a := range *artifacts {
		newExpectedArtifact := client.NewExpectedArtifact()
		if a.DefaultArtifact != nil && len(a.DefaultArtifact) > 0 {
			newExpectedArtifact.DefaultArtifact = client.ExpectedArtifactDefaultArtifact(a.DefaultArtifact[0])
		}
		newExpectedArtifact.DisplayName = a.DisplayName
		newExpectedArtifact.ID = a.ID
		if a.MatchArtifact != nil && len(a.MatchArtifact) > 0 {
			newExpectedArtifact.MatchArtifact = client.ExpectedArtifactMatchArtifact(a.MatchArtifact[0])
		}
		newExpectedArtifact.UseDefaultArtifact = a.UseDefaultArtifact
		newExpectedArtifact.UsePriorArtifact = a.UsePriorArtifact
		newList = append(newList, newExpectedArtifact)
	}
	return &newList
}

func fromClientExpectedArtifacts(artifacts *[]*client.ExpectedArtifact) *[]*expectedArtifact {
	if artifacts == nil {
		return nil
	}
	newList := []*expectedArtifact{}
	for _, a := range *artifacts {
		newExpectedArtifact := expectedArtifact{
			DefaultArtifact:    []expectedArtifactDefaultArtifact{expectedArtifactDefaultArtifact(a.DefaultArtifact)},
			DisplayName:        a.DisplayName,
			ID:                 a.ID,
			MatchArtifact:      []expectedArtifactMatchArtifact{expectedArtifactMatchArtifact(a.MatchArtifact)},
			UseDefaultArtifact: a.UseDefaultArtifact,
			UsePriorArtifact:   a.UsePriorArtifact,
		}
		newList = append(newList, &newExpectedArtifact)
	}
	return &newList
}
