package provider

import (
	"github.com/get-bridge/terraform-provider-spinnaker/client"
	"github.com/hashicorp/terraform/helper/schema"
)

type runJobManifestStage struct {
	baseStage `mapstructure:",squash"`

	Account               string   `mapstructure:"account"`
	Application           string   `mapstructure:"application"`
	CloudProvider         string   `mapstructure:"cloud_provider"`
	ConsumeArtifactSource string   `mapstructure:"consume_artifact_source"`
	Credentials           string   `mapstructure:"credentials"`
	Manifest              manifest `mapstructure:"manifest"`
	PropertyFile          string   `mapstructure:"property_file"`
	Source                string   `mapstructure:"source"`
}

func newRunJobManifestStage() *runJobManifestStage {
	return &runJobManifestStage{
		baseStage: *newBaseStage(),
	}
}

func (s *runJobManifestStage) toClientStage(config *client.Config, refID string) (client.Stage, error) {
	cs := client.NewRunJobManifestStage()
	err := s.baseToClientStage(&cs.BaseStage, refID, newDefaultNotificationInterface)
	if err != nil {
		return nil, err
	}

	cs.Account = s.Account
	cs.Application = s.Application
	cs.CloudProvider = s.CloudProvider
	cs.ConsumeArtifactSource = s.ConsumeArtifactSource
	cs.Credentials = s.Credentials
	cs.Manifest = client.Manifest(s.Manifest)
	cs.PropertyFile = s.PropertyFile
	cs.Source = s.Source

	return cs, nil
}

func (*runJobManifestStage) fromClientStage(cs client.Stage) (stage, error) {
	clientStage := cs.(*client.RunJobManifestStage)
	newStage := newRunJobManifestStage()
	err := newStage.baseFromClientStage(&clientStage.BaseStage, newDefaultNotificationInterface)
	if err != nil {
		return nil, err
	}

	newStage.Account = clientStage.Account
	newStage.Application = clientStage.Application
	newStage.CloudProvider = clientStage.CloudProvider
	newStage.ConsumeArtifactSource = clientStage.ConsumeArtifactSource
	newStage.Credentials = clientStage.Credentials
	newStage.Manifest = manifest(clientStage.Manifest)
	newStage.PropertyFile = clientStage.PropertyFile
	newStage.Source = clientStage.Source

	return newStage, nil
}

func (s *runJobManifestStage) SetResourceData(d *schema.ResourceData) error {
	err := s.baseSetResourceData(d)
	if err != nil {
		return err
	}

	err = d.Set("account", s.Account)
	if err != nil {
		return err
	}
	err = d.Set("application", s.Application)
	if err != nil {
		return err
	}
	err = d.Set("cloud_provider", s.CloudProvider)
	if err != nil {
		return err
	}
	err = d.Set("consume_artifact_source", s.ConsumeArtifactSource)
	if err != nil {
		return err
	}
	err = d.Set("credentials", s.Credentials)
	if err != nil {
		return err
	}
	err = d.Set("manifest", s.Manifest)
	if err != nil {
		return err
	}
	err = d.Set("property_file", s.PropertyFile)
	if err != nil {
		return err
	}
	err = d.Set("source", s.Source)
	if err != nil {
		return err
	}
	return nil
}
