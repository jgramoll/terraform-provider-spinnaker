package provider

import (
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

type deployManifests []string

func newDeployManifests() *deployManifests {
	return &deployManifests{}
}

func (manifests *deployManifests) toClientDeployManifests() *client.DeployManifests {
	newManifests := client.NewDeployManifests()
	for _, m := range *manifests {
		*newManifests = append(*newManifests, m)
	}
	return newManifests
}

func fromClientDeployManifests(clientManifests *client.DeployManifests) *deployManifests {
	if clientManifests == nil {
		return nil
	}
	newManifests := newDeployManifests()
	for _, m := range *clientManifests {
		*newManifests = append(*newManifests, m)
	}
	return newManifests
}
