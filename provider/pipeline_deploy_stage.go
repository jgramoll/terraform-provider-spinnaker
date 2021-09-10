package provider

import (
	"github.com/get-bridge/terraform-provider-spinnaker/client"
	"github.com/hashicorp/terraform/helper/schema"
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
	err := s.baseToClientStage(&cs.BaseStage, refID, newDefaultNotificationInterface)
	if err != nil {
		return nil, err
	}

	if s.Clusters != nil {
		cs.Clusters = s.Clusters.toClientClusters()
	}

	return cs, nil
}

func (*deployStage) fromClientStage(cs client.Stage) (stage, error) {
	clientStage := cs.(*client.DeployStage)
	newStage := newDeployStage()
	err := newStage.baseFromClientStage(&clientStage.BaseStage, newDefaultNotificationInterface)
	if err != nil {
		return nil, err
	}

	newStage.Clusters = newStage.Clusters.fromClientClusters(clientStage.Clusters)

	return newStage, nil
}

func (s *deployStage) SetResourceData(d *schema.ResourceData) error {
	err := s.baseSetResourceData(d)
	if err != nil {
		return err
	}

	return d.Set("cluster", s.Clusters)
}
