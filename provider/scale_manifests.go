package provider

import (
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

type scaleManifests []string

func newScaleManifests() *scaleManifests {
	return &scaleManifests{}
}

func (manifests *scaleManifests) toClientScaleManifests() *client.ScaleManifests {
	newManifests := client.NewScaleManifests()
	for _, m := range *manifests {
		*newManifests = append(*newManifests, m)
	}
	return newManifests
}

func fromClientScaleManifests(clientManifests *client.ScaleManifests) *scaleManifests {
	if clientManifests == nil {
		return nil
	}
	newManifests := newScaleManifests()
	for _, m := range *clientManifests {
		*newManifests = append(*newManifests, m)
	}
	return newManifests
}
