package provider

import (
	"bytes"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
	"gopkg.in/yaml.v2"
	kyaml "k8s.io/apimachinery/pkg/util/yaml"
)

type patchManifestOptions struct {
	MergeStrategy string `mapstructure:"merge_strategy"`
	Record        bool   `mapstructure:"record"`
}

type patchManifestStage struct {
	baseStage `mapstructure:",squash"`

	Account       string `mapstructure:"account"`
	App           string `mapstructure:"app"`
	CloudProvider string `mapstructure:"cloud_provider"`
	Cluster       string `mapstructure:"cluster"`
	Criteria      string `mapstructure:"criteria"`
	Kind          string `mapstructure:"kind"`
	// kinds string `json:"kinds"`
	// labelSelectors string `json:"labelSelectors"`
	Location     string                 `mapstructure:"location"`
	ManifestName string                 `mapstructure:"manifest_name"`
	Mode         string                 `mapstructure:"mode"`
	Options      []patchManifestOptions `mapstructure:"options"`
	PatchBody    []string               `mapstructure:"patch_body"`
	Source       string                 `mapstructure:"source"`
}

func newPatchManifestStage() *patchManifestStage {
	return &patchManifestStage{
		baseStage: *newBaseStage(),
	}
}

func (s *patchManifestStage) toClientStage(config *client.Config, refID string) (client.Stage, error) {
	cs := client.NewPatchManifestStage()
	err := s.baseToClientStage(&cs.BaseStage, refID, newDefaultNotificationInterface)
	if err != nil {
		return nil, err
	}

	cs.Account = s.Account
	cs.App = s.App
	cs.CloudProvider = s.CloudProvider
	cs.Cluster = s.Cluster
	cs.Criteria = s.Criteria
	cs.Kind = s.Kind
	cs.Location = s.Location
	cs.ManifestName = s.ManifestName
	cs.Mode = s.Mode
	if len(s.Options) > 0 {
		cs.Options = client.PatchManifestOptions(s.Options[0])
	}
	cs.PatchBody = []map[string]interface{}{}
	for _, ps := range s.PatchBody {
		var p map[string]interface{}
		if err := kyaml.NewYAMLOrJSONDecoder(bytes.NewReader([]byte(ps)), 50).Decode(&p); err != nil {
			return nil, err
		}
		cs.PatchBody = append(cs.PatchBody, p)
	}

	cs.Source = s.Source

	return cs, nil
}

func (*patchManifestStage) fromClientStage(cs client.Stage) (stage, error) {
	clientStage := cs.(*client.PatchManifestStage)
	newStage := newPatchManifestStage()
	err := newStage.baseFromClientStage(&clientStage.BaseStage, newDefaultNotificationInterface)
	if err != nil {
		return nil, err
	}

	newStage.Account = clientStage.Account
	newStage.App = clientStage.App
	newStage.CloudProvider = clientStage.CloudProvider
	newStage.Cluster = clientStage.Cluster
	newStage.Criteria = clientStage.Criteria
	newStage.Kind = clientStage.Kind
	newStage.Location = clientStage.Location
	newStage.ManifestName = clientStage.ManifestName
	newStage.Mode = clientStage.Mode
	newStage.Options = []patchManifestOptions{patchManifestOptions(clientStage.Options)}
	newStage.PatchBody = []string{}
	for _, ps := range clientStage.PatchBody {
		p := bytes.NewBuffer(nil)
		if err := yaml.NewEncoder(p).Encode(ps); err != nil {
			return nil, err
		}
		newStage.PatchBody = append(newStage.PatchBody, p.String())
	}
	newStage.Source = clientStage.Source

	return newStage, nil
}

func (s *patchManifestStage) SetResourceData(d *schema.ResourceData) error {
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
	err = d.Set("options", s.Options)
	if err != nil {
		return err
	}
	err = d.Set("patch_body", s.PatchBody)
	if err != nil {
		return err
	}
	err = d.Set("source", s.Source)
	if err != nil {
		return err
	}
	return nil
}
