package provider

import (
	"github.com/get-bridge/terraform-provider-spinnaker/client"
	"github.com/hashicorp/terraform/helper/schema"
)

type enableServerGroupStage struct {
	baseStage `mapstructure:",squash"`

	CloudProvider                  string   `mapstructure:"cloud_provider"`
	CloudProviderType              string   `mapstructure:"cloud_provider_type"`
	Cluster                        string   `mapstructure:"cluster"`
	Credentials                    string   `mapstructure:"credentials"`
	InterestingHealthProviderNames []string `mapstructure:"interesting_health_provider_names"`
	Namespaces                     []string `mapstructure:"namespaces"`
	Regions                        []string `mapstructure:"regions"`
	Target                         string   `mapstructure:"target"`
}

func newEnableServerGroupStage() *enableServerGroupStage {
	return &enableServerGroupStage{
		baseStage: *newBaseStage(),
	}
}

func (s *enableServerGroupStage) toClientStage(config *client.Config, refID string) (client.Stage, error) {
	cs := client.NewEnableServerGroupStage()
	err := s.baseToClientStage(&cs.BaseStage, refID, newDefaultNotificationInterface)
	if err != nil {
		return nil, err
	}

	cs.CloudProvider = s.CloudProvider
	cs.CloudProviderType = s.CloudProviderType
	cs.Cluster = s.Cluster
	cs.Credentials = s.Credentials
	cs.InterestingHealthProviderNames = s.InterestingHealthProviderNames
	cs.Namespaces = s.Namespaces
	cs.Regions = s.Regions
	cs.Target = s.Target

	return cs, nil
}

func (*enableServerGroupStage) fromClientStage(cs client.Stage) (stage, error) {
	clientStage := cs.(*client.EnableServerGroupStage)
	newStage := newEnableServerGroupStage()
	err := newStage.baseFromClientStage(&clientStage.BaseStage, newDefaultNotificationInterface)
	if err != nil {
		return nil, err
	}

	newStage.CloudProvider = clientStage.CloudProvider
	newStage.CloudProviderType = clientStage.CloudProviderType
	newStage.Cluster = clientStage.Cluster
	newStage.Credentials = clientStage.Credentials
	newStage.InterestingHealthProviderNames = clientStage.InterestingHealthProviderNames
	newStage.Namespaces = clientStage.Namespaces
	newStage.Regions = clientStage.Regions
	newStage.Target = clientStage.Target

	return newStage, nil
}

func (s *enableServerGroupStage) SetResourceData(d *schema.ResourceData) error {
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
	err = d.Set("interesting_health_provider_names", s.InterestingHealthProviderNames)
	if err != nil {
		return err
	}
	err = d.Set("namespaces", s.Namespaces)
	if err != nil {
		return err
	}
	err = d.Set("regions", s.Regions)
	if err != nil {
		return err
	}
	err = d.Set("target", s.Target)
	if err != nil {
		return err
	}

	return nil
}
