package provider

import (
	"github.com/get-bridge/terraform-provider-spinnaker/client"
	"github.com/hashicorp/terraform/helper/schema"
)

type scaleManifestStage struct {
	baseStage `mapstructure:",squash"`

	Account        string                 `mapstructure:"account"`
	Application    string                 `mapstructure:"application"`
	CloudProvider  string                 `mapstructure:"cloud_provider"`
	Cluster        string                 `mapstructure:"cluster"`
	Criteria       string                 `mapstructure:"criteria"`
	Kind           string                 `mapstructure:"kind"`
	Kinds          []string               `mapstructure:"kinds"`
	LabelSelectors map[string]interface{} `mapstructure:"label_selectors"`
	Location       string                 `mapstructure:"location"`
	ManifestName   string                 `mapstructure:"manifest_name"`
	Mode           string                 `mapstructure:"mode"`
	Replicas       string                 `mapstructure:"replicas"`
}

func newScaleManifestStage() *scaleManifestStage {
	return &scaleManifestStage{
		baseStage: *newBaseStage(),
		// TODO defaults
	}
}

func (s *scaleManifestStage) toClientStage(config *client.Config, refID string) (client.Stage, error) {
	cs := client.NewScaleManifestStage()
	err := s.baseToClientStage(&cs.BaseStage, refID, newDefaultNotificationInterface)
	if err != nil {
		return nil, err
	}

	cs.Account = s.Account
	cs.Application = s.Application
	cs.CloudProvider = s.CloudProvider
	cs.Cluster = s.Cluster
	cs.Criteria = s.Criteria
	cs.Kind = s.Kind
	cs.Kinds = s.Kinds
	cs.LabelSelectors = s.LabelSelectors
	cs.Location = s.Location
	cs.ManifestName = s.ManifestName
	cs.Mode = s.Mode
	cs.Replicas = s.Replicas
	return cs, nil
}

func (*scaleManifestStage) fromClientStage(cs client.Stage) (stage, error) {
	clientStage := cs.(*client.ScaleManifestStage)
	newStage := newScaleManifestStage()
	err := newStage.baseFromClientStage(&clientStage.BaseStage, newDefaultNotificationInterface)
	if err != nil {
		return nil, err
	}

	newStage.Account = clientStage.Account
	newStage.Application = clientStage.Application
	newStage.CloudProvider = clientStage.CloudProvider
	newStage.Cluster = clientStage.Cluster
	newStage.Criteria = clientStage.Criteria
	newStage.Kind = clientStage.Kind
	newStage.Kinds = clientStage.Kinds
	newStage.LabelSelectors = clientStage.LabelSelectors
	newStage.Location = clientStage.Location
	newStage.ManifestName = clientStage.ManifestName
	newStage.Mode = clientStage.Mode
	newStage.Replicas = clientStage.Replicas

	return newStage, nil
}

func (s *scaleManifestStage) SetResourceData(d *schema.ResourceData) error {
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
	err = d.Set("cluster", s.Cluster)
	if err != nil {
		return err
	}
	err = d.Set("criteria", s.Criteria)
	if err != nil {
		return err
	}
	err = d.Set("kind", s.Kind)
	if err != nil {
		return err
	}
	err = d.Set("kinds", s.Kinds)
	if err != nil {
		return err
	}
	err = d.Set("label_selectors", s.LabelSelectors)
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
	return d.Set("replicas", s.Replicas)
}
