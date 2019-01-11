package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

type rollbackClusterStage struct {
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

	TargetHealthyRollbackPercentage int `mapstructure:"target_healthy_rollback_percentage"`
}

func newRollbackClusterStage() *rollbackClusterStage {
	return &rollbackClusterStage{Type: client.RollbackClusterType}
}

func (s *rollbackClusterStage) toClientStage() (client.Stage, error) {
	cs := client.NewRollbackClusterStage()
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
	cs.TargetHealthyRollbackPercentage = s.TargetHealthyRollbackPercentage

	if len(s.Moniker) > 0 {
		newMoniker := client.Moniker(s.Moniker[0])
		cs.Moniker = &newMoniker
	}

	return cs, nil
}

// TODO can we just update the ptr?
func (s *rollbackClusterStage) fromClientStage(cs client.Stage) stage {
	rollbackStage := cs.(*client.RollbackClusterStage)
	newStage := newRollbackClusterStage()
	newStage.Name = rollbackStage.Name
	newStage.RefID = rollbackStage.RefID
	newStage.RequisiteStageRefIds = rollbackStage.RequisiteStageRefIds

	if rollbackStage.StageEnabled != nil {
		newStage.StageEnabled = append(newStage.StageEnabled, stageEnabled(*rollbackStage.StageEnabled))
	}

	newStage.CloudProvider = rollbackStage.CloudProvider
	newStage.CloudProviderType = rollbackStage.CloudProviderType
	newStage.Cluster = rollbackStage.Cluster
	newStage.Credentials = rollbackStage.Credentials
	newStage.Regions = rollbackStage.Regions
	newStage.TargetHealthyRollbackPercentage = rollbackStage.TargetHealthyRollbackPercentage

	if rollbackStage.Moniker != nil {
		newStage.Moniker = append(newStage.Moniker, moniker(*rollbackStage.Moniker))
	}

	return newStage
}

func (s *rollbackClusterStage) SetResourceData(d *schema.ResourceData) {
	// TODO
	d.Set("name", s.Name)
}

func (s *rollbackClusterStage) SetRefID(id string) {
	s.RefID = id
}

func (s *rollbackClusterStage) GetRefID() string {
	return s.RefID
}
