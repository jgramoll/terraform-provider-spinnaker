package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

type deployStage struct {
	baseStage `mapstructure:",squash"`

	Clusters *deployStageClusters `mapstructure:"cluster"`
}

func newDeployStage() *deployStage {
	return &deployStage{
		baseStage: *newBaseStage(),
	}
}

func (s *deployStage) toClientStage(config *client.Config, refID string) (client.Stage, error) {
	cs := client.NewDeployStage()
	err := s.baseToClientStage(&cs.BaseStage, refID)
	if err != nil {
		return nil, err
	}

	if s.Clusters != nil {
		cs.Clusters = s.Clusters.toClientClusters()
	}

	return cs, nil
}

func (*deployStage) fromClientStage(cs client.Stage) stage {
	clientStage := cs.(*client.DeployStage)
	newStage := newDeployStage()
	newStage.baseFromClientStage(&clientStage.BaseStage)

	newStage.Clusters = newStage.Clusters.fromClientClusters(clientStage.Clusters)

	return newStage
}

func (s *deployStage) SetResourceData(d *schema.ResourceData) error {
	err := s.baseSetResourceData(d)
	if err != nil {
		return err
	}

	return d.Set("cluster", s.Clusters)
}
