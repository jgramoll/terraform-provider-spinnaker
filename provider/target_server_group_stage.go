package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

type targetServerGroupStage struct {
	CloudProvider     string      `mapstructure:"cloud_provider"`
	CloudProviderType string      `mapstructure:"cloud_provider_type"`
	Cluster           string      `mapstructure:"cluster"`
	Credentials       string      `mapstructure:"credentials"`
	Moniker           *[]*moniker `mapstructure:"moniker"`
	Regions           []string    `mapstructure:"regions"`
	Target            string      `mapstructure:"target"`
}

func (s *targetServerGroupStage) targetServerGroupStageToClient(cs *client.TargetServerGroupStage) error {
	cs.CloudProvider = s.CloudProvider
	cs.CloudProviderType = s.CloudProviderType
	cs.Cluster = s.Cluster
	cs.Credentials = s.Credentials
	cs.Moniker = toClientMoniker(s.Moniker)
	cs.Regions = s.Regions
	cs.Target = s.Target

	return nil
}

func (s *targetServerGroupStage) targetServerGroupStageFromClientStage(clientStage *client.TargetServerGroupStage) error {
	s.CloudProvider = clientStage.CloudProvider
	s.CloudProviderType = clientStage.CloudProviderType
	s.Cluster = clientStage.Cluster
	s.Credentials = clientStage.Credentials
	s.Moniker = fromClientMoniker(clientStage.Moniker)
	s.Regions = clientStage.Regions
	s.Target = clientStage.Target

	return nil
}

func (s *targetServerGroupStage) targetServerGroupSetResourceData(d *schema.ResourceData) error {
	var err error
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
	return d.Set("target", s.Target)
}
