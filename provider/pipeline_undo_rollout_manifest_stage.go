package provider

import (
	"github.com/get-bridge/terraform-provider-spinnaker/client"
	"github.com/hashicorp/terraform/helper/schema"
)

type undoRolloutManifestStage struct {
	baseStage `mapstructure:",squash"`

	Account          string `mapstructure:"account"`
	CloudProvider    string `mapstructure:"cloud_provider"`
	Location         string `mapstructure:"location"`
	ManifestName     string `mapstructure:"manifest_name"`
	Mode             string `mapstructure:"mode"`
	NumRevisionsBack int    `mapstructure:"num_revisions_back"`
}

func newUndoRolloutManifestStage() *undoRolloutManifestStage {
	return &undoRolloutManifestStage{
		baseStage: *newBaseStage(),
	}
}

func (s *undoRolloutManifestStage) toClientStage(config *client.Config, refID string) (client.Stage, error) {
	cs := client.NewUndoRolloutManifestStage()
	err := s.baseToClientStage(&cs.BaseStage, refID, newDefaultNotificationInterface)
	if err != nil {
		return nil, err
	}

	cs.Account = s.Account
	cs.CloudProvider = s.CloudProvider
	cs.Location = s.Location
	cs.ManifestName = s.ManifestName
	cs.Mode = s.Mode
	cs.NumRevisionsBack = s.NumRevisionsBack

	return cs, nil
}

func (*undoRolloutManifestStage) fromClientStage(cs client.Stage) (stage, error) {
	clientStage := cs.(*client.UndoRolloutManifestStage)
	newStage := newUndoRolloutManifestStage()
	err := newStage.baseFromClientStage(&clientStage.BaseStage, newDefaultNotificationInterface)
	if err != nil {
		return nil, err
	}

	newStage.Account = clientStage.Account
	newStage.CloudProvider = clientStage.CloudProvider
	newStage.Location = clientStage.Location
	newStage.ManifestName = clientStage.ManifestName
	newStage.Mode = clientStage.Mode
	newStage.NumRevisionsBack = clientStage.NumRevisionsBack

	return newStage, nil
}

func (s *undoRolloutManifestStage) SetResourceData(d *schema.ResourceData) error {
	err := s.baseSetResourceData(d)
	if err != nil {
		return err
	}

	err = d.Set("account", s.Account)
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
	err = d.Set("num_revisions_back", s.NumRevisionsBack)
	if err != nil {
		return err
	}

	return nil
}
