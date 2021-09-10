package provider

import (
	"github.com/get-bridge/terraform-provider-spinnaker/client"
	"github.com/hashicorp/terraform/helper/schema"
)

type resizeServerGroupStage struct {
	baseStage `mapstructure:",squash"`

	Action            string       `mapstructure:"action"`
	Capacity          *[]*capacity `mapstructure:"capacity"`
	CloudProvider     string       `mapstructure:"cloud_provider"`
	CloudProviderType string       `mapstructure:"cloud_provider_type"`
	Cluster           string       `mapstructure:"cluster"`
	Credentials       string       `mapstructure:"credentials"`
	Moniker           *[]*moniker  `mapstructure:"moniker"`
	Regions           []string     `mapstructure:"regions"`
	ResizeType        string       `mapstructure:"resize_type"`
	Target            string       `mapstructure:"target"`

	TargetHealthyDeployPercentage int `mapstructure:"target_healthy_deploy_percentage"`
}

func newResizeServerGroupStage() *resizeServerGroupStage {
	return &resizeServerGroupStage{
		baseStage: *newBaseStage(),
	}
}

func (s *resizeServerGroupStage) toClientStage(config *client.Config, refID string) (client.Stage, error) {
	cs := client.NewResizeServerGroupStage()
	err := s.baseToClientStage(&cs.BaseStage, refID, newDefaultNotificationInterface)
	if err != nil {
		return nil, err
	}

	cs.Action = s.Action
	cs.Capacity = toClientCapacity(s.Capacity)
	cs.CloudProvider = s.CloudProvider
	cs.CloudProviderType = s.CloudProviderType
	cs.Cluster = s.Cluster
	cs.Credentials = s.Credentials
	cs.Moniker = toClientMoniker(s.Moniker)
	cs.Regions = s.Regions
	cs.ResizeType = s.ResizeType
	cs.Target = s.Target
	cs.TargetHealthyDeployPercentage = s.TargetHealthyDeployPercentage

	return cs, nil
}

func (*resizeServerGroupStage) fromClientStage(cs client.Stage) (stage, error) {
	clientStage := cs.(*client.ResizeServerGroupStage)
	newStage := newResizeServerGroupStage()
	err := newStage.baseFromClientStage(&clientStage.BaseStage, newDefaultNotificationInterface)
	if err != nil {
		return nil, err
	}

	newStage.Action = clientStage.Action
	newStage.Capacity = fromClientCapacity(clientStage.Capacity)
	newStage.CloudProvider = clientStage.CloudProvider
	newStage.CloudProviderType = clientStage.CloudProviderType
	newStage.Cluster = clientStage.Cluster
	newStage.Credentials = clientStage.Credentials
	newStage.Moniker = fromClientMoniker(clientStage.Moniker)
	newStage.Regions = clientStage.Regions
	newStage.ResizeType = clientStage.ResizeType
	newStage.Target = clientStage.Target
	newStage.TargetHealthyDeployPercentage = clientStage.TargetHealthyDeployPercentage

	return newStage, nil
}

func (s *resizeServerGroupStage) SetResourceData(d *schema.ResourceData) error {
	err := s.baseSetResourceData(d)
	if err != nil {
		return err
	}

	err = d.Set("action", s.Action)
	if err != nil {
		return err
	}
	err = d.Set("capacity", s.Capacity)
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
	err = d.Set("resize_type", s.ResizeType)
	if err != nil {
		return err
	}
	err = d.Set("target", s.Target)
	if err != nil {
		return err
	}
	return d.Set("target_healthy_deploy_percentage", s.TargetHealthyDeployPercentage)
}
