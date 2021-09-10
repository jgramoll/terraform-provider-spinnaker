package provider

import (
	"github.com/get-bridge/terraform-provider-spinnaker/client"
	"github.com/hashicorp/terraform/helper/schema"
)

type deleteManifestStage struct {
	baseStage `mapstructure:",squash"`

	Account       string                    `mapstructure:"account"`
	App           string                    `mapstructure:"app"`
	CloudProvider string                    `mapstructure:"cloud_provider"`
	Location      string                    `mapstructure:"location"`
	ManifestName  string                    `mapstructure:"manifest_name"`
	Mode          string                    `mapstructure:"mode"`
	Options       *[]*deleteManifestOptions `mapstructure:"options"`
}

func newDeleteManifestStage() *deleteManifestStage {
	return &deleteManifestStage{
		baseStage: *newBaseStage(),
		Options:   &[]*deleteManifestOptions{},
	}
}

func (s *deleteManifestStage) toClientStage(config *client.Config, refID string) (client.Stage, error) {
	cs := client.NewDeleteManifestStage()
	err := s.baseToClientStage(&cs.BaseStage, refID, newDefaultNotificationInterface)
	if err != nil {
		return nil, err
	}

	cs.Account = s.Account
	cs.App = s.App
	cs.CloudProvider = s.CloudProvider
	cs.Location = s.Location
	cs.ManifestName = s.ManifestName
	mode, err := client.ParseDeleteManifestMode(s.Mode)
	if err != nil {
		return nil, err
	}
	cs.Mode = mode
	cs.Options = toClientOptions(s.Options)

	return cs, nil
}

func (*deleteManifestStage) fromClientStage(cs client.Stage) (stage, error) {
	clientStage := cs.(*client.DeleteManifestStage)
	newStage := newDeleteManifestStage()
	err := newStage.baseFromClientStage(&clientStage.BaseStage, newDefaultNotificationInterface)
	if err != nil {
		return nil, err
	}

	newStage.Account = clientStage.Account
	newStage.App = clientStage.App
	newStage.CloudProvider = clientStage.CloudProvider
	newStage.Location = clientStage.Location
	newStage.ManifestName = clientStage.ManifestName
	newStage.Mode = clientStage.Mode.String()
	newStage.Options = fromClientDeleteManifestOptions(clientStage.Options)

	return newStage, nil
}

func (s *deleteManifestStage) SetResourceData(d *schema.ResourceData) error {
	err := s.baseSetResourceData(d)
	if err != nil {
		return err
	}

	err = d.Set("account", s.Account)
	if err != nil {
		return err
	}
	err = d.Set("app", s.App)
	if err != nil {
		return err
	}
	err = d.Set("cloud_provider", s.CloudProvider)
	if err != nil {
		return err
	}
	err = d.Set("location", s.Location)
	if err != nil {
		return err
	}
	err = d.Set("manifest_name", s.ManifestName)
	if err != nil {
		return err
	}
	err = d.Set("mode", s.Mode)
	if err != nil {
		return err
	}
	return d.Set("options", s.Options)
}
