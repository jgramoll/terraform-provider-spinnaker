package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

type destroyServerGroupStage struct {
	// TODO does this baseStage work?
	// baseStage
	Name                 string           `mapstructure:"name"`
	RefID                string           `mapstructure:"ref_id"`
	Type                 client.StageType `mapstructure:"type"`
	RequisiteStageRefIds []string         `mapstructure:"requisite_stage_ref_ids"`
	StageEnabled         []stageEnabled   `mapstructure:"stage_enabled"`

	CloudProvider     string    `mapstructure:"cloud_provider"`
	CloudProviderType string    `mapstructure:"cloud_provider_type"`
	Cluster           string    `mapstructure:"cluster"`
	Credentials       string    `mapstructure:"credentials"`
	Moniker           []moniker `mapstructure:"moniker"`
	Regions           []string  `mapstructure:"regions"`
	Target            string    `mapstructure:"target"`
}

func newDestroyServerGroupStage() *destroyServerGroupStage {
	return &destroyServerGroupStage{Type: client.DestroyServerGroupType}
}

func (s *destroyServerGroupStage) toClientStage() client.Stage {
	cs := client.NewDestroyServerGroupStage()
	cs.Name = s.Name
	cs.RefID = s.RefID
	cs.RequisiteStageRefIds = s.RequisiteStageRefIds

	if len(s.StageEnabled) > 0 {
		newStageEnabled := client.StageEnabled(s.StageEnabled[0])
		cs.StageEnabled = &newStageEnabled
	}

	cs.CloudProvider = s.CloudProvider
	cs.CloudProviderType = s.CloudProviderType
	cs.Cluster = s.Cluster
	cs.Credentials = s.Credentials
	cs.Regions = s.Regions
	cs.Target = s.Target

	if len(s.Moniker) > 0 {
		newMoniker := client.Moniker(s.Moniker[0])
		cs.Moniker = &newMoniker
	}

	return cs
}

// TODO can we just update the ptr?
func (s *destroyServerGroupStage) fromClientStage(cs client.Stage) stage {
	clientStage := cs.(*client.DestroyServerGroupStage)
	newStage := newDestroyServerGroupStage()
	newStage.Name = clientStage.Name
	newStage.RefID = clientStage.RefID
	newStage.RequisiteStageRefIds = clientStage.RequisiteStageRefIds

	if clientStage.StageEnabled != nil {
		newStage.StageEnabled = append(newStage.StageEnabled, stageEnabled(*clientStage.StageEnabled))
	}

	newStage.CloudProvider = clientStage.CloudProvider
	newStage.CloudProviderType = clientStage.CloudProviderType
	newStage.Cluster = clientStage.Cluster
	newStage.Credentials = clientStage.Credentials
	newStage.Regions = clientStage.Regions
	newStage.Target = clientStage.Target

	if clientStage.Moniker != nil {
		newStage.Moniker = append(newStage.Moniker, moniker(*clientStage.Moniker))
	}

	return newStage
}

func (s *destroyServerGroupStage) SetResourceData(d *schema.ResourceData) {
	// TODO
	d.Set("name", s.Name)
}

func (s *destroyServerGroupStage) SetRefID(id string) {
	s.RefID = id
}

func (s *destroyServerGroupStage) GetRefID() string {
	return s.RefID
}
