package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

type findImageFromTagsStage struct {
	baseStage `mapstructure:",squash"`

	CloudProvider     string            `mapstructure:"cloud_provider"`
	CloudProviderType string            `mapstructure:"cloud_provider_type"`
	PackageName       string            `mapstructure:"package_name"`
	Regions           []string          `mapstructure:"regions"`
	Tags              map[string]string `mapstructure:"tags"`
}

func newFindImageFromTagsStage() *findImageFromTagsStage {
	return &findImageFromTagsStage{
		baseStage: *newBaseStage(),
	}
}

func (s *findImageFromTagsStage) toClientStage(config *client.Config, refID string) (client.Stage, error) {
	cs := client.NewFindImageStage()
	err := s.baseToClientStage(&cs.BaseStage, refID)
	if err != nil {
		return nil, err
	}

	cs.CloudProvider = s.CloudProvider
	cs.CloudProviderType = s.CloudProviderType
	cs.PackageName = s.PackageName
	cs.Regions = s.Regions
	cs.Tags = s.Tags

	return cs, nil
}

func (*findImageFromTagsStage) fromClientStage(cs client.Stage) stage {
	clientStage := cs.(*client.FindImageFromTagsStage)
	newStage := newFindImageFromTagsStage()
	newStage.baseFromClientStage(&clientStage.BaseStage)

	newStage.CloudProvider = clientStage.CloudProvider
	newStage.CloudProviderType = clientStage.CloudProviderType
	newStage.PackageName = clientStage.PackageName
	newStage.Regions = clientStage.Regions
	newStage.Tags = clientStage.Tags

	return newStage
}

func (s *findImageFromTagsStage) SetResourceData(d *schema.ResourceData) error {
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
	err = d.Set("package_name", s.PackageName)
	if err != nil {
		return err
	}
	err = d.Set("regions", s.Regions)
	if err != nil {
		return err
	}
	return d.Set("tags", s.Tags)

}
