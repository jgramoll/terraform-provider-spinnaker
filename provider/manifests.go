package provider

import (
	"github.com/get-bridge/terraform-provider-spinnaker/client"
)

type manifest string
type manifests []manifest

func newManifests() *manifests {
	return &manifests{}
}

func (manifestArray *manifests) toClientManifests() *client.Manifests {
	newManifests := client.NewManifests()
	for _, m := range *manifestArray {
		*newManifests = append(*newManifests, client.Manifest(m))
	}
	return newManifests
}

func fromClientManifests(clientManifests *client.Manifests) *manifests {
	if clientManifests == nil {
		return nil
	}
	newManifests := newManifests()
	for _, m := range *clientManifests {
		*newManifests = append(*newManifests, manifest(m))
	}
	return newManifests
}
