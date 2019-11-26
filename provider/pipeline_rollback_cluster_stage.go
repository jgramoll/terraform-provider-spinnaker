package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

type rollbackClusterStage struct {
	baseStage `mapstructure:",squash"`

	CloudProvider     string      `mapstructure:"cloud_provider"`
	CloudProviderType string      `mapstructure:"cloud_provider_type"`
	Cluster           string      `mapstructure:"cluster"`
	Credentials       string      `mapstructure:"credentials"`
	Moniker           *[]*moniker `mapstructure:"moniker"`
	Regions           []string    `mapstructure:"regions"`

	TargetHealthyRollbackPercentage int `mapstructure:"target_healthy_rollback_percentage"`
}

func newRollbackClusterStage() *rollbackClusterStage {
	return &rollbackClusterStage{
		baseStage: *newBaseStage(),
	}
}

func (s *rollbackClusterStage) toClientStage(config *client.Config, refId string) (client.Stage, error) {
	cs := client.NewRollbackClusterStage()
	err := s.baseToClientStage(&cs.BaseStage, refId)
	if err != nil {
		return nil, err
	}

	cs.CloudProvider = s.CloudProvider
	cs.CloudProviderType = s.CloudProviderType
	cs.Cluster = s.Cluster
	cs.Credentials = s.Credentials
	cs.Moniker = toClientMoniker(s.Moniker)
	cs.Regions = s.Regions
	cs.TargetHealthyRollbackPercentage = s.TargetHealthyRollbackPercentage

	return cs, nil
}

func (s *rollbackClusterStage) fromClientStage(cs client.Stage) stage {
	clientStage := cs.(*client.RollbackClusterStage)
	newStage := newRollbackClusterStage()
	newStage.baseFromClientStage(&clientStage.BaseStage)

	newStage.CloudProvider = clientStage.CloudProvider
	newStage.CloudProviderType = clientStage.CloudProviderType
	newStage.Cluster = clientStage.Cluster
	newStage.Credentials = clientStage.Credentials
	newStage.Moniker = fromClientMoniker(clientStage.Moniker)
	newStage.Regions = clientStage.Regions
	newStage.TargetHealthyRollbackPercentage = clientStage.TargetHealthyRollbackPercentage

	return newStage
}

func (s *rollbackClusterStage) SetResourceData(d *schema.ResourceData) error {
	err := s.baseSetResourceData(d)
	if err != nil {
		return err
	}

	err = d.Set("cloud_provider", s.CloudProvider)
	if err != nil {
		return err
	}
	err = d.Set("cloud_provider_type", s.CloudProviderType)
	if err != nil {
		return err
	}
	err = d.Set("cluster", s.Cluster)
	if err != nil {
		return err
	}
	err = d.Set("credentials", s.Credentials)
	if err != nil {
		return err
	}
	err = d.Set("moniker", s.Moniker)
	if err != nil {
		return err
	}
	err = d.Set("regions", s.Regions)
	if err != nil {
		return err
	}
	return d.Set("target_healthy_rollback_percentage", s.TargetHealthyRollbackPercentage)
}
