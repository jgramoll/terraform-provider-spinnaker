package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

type deployStage struct {
	Name                 string           `mapstructure:"name"`
	RefID                string           `mapstructure:"ref_id"`
	Type                 client.StageType `mapstructure:"type"`
	RequisiteStageRefIds []string         `mapstructure:"requisite_stage_ref_ids"`

	CompleteOtherBranchesThenFail bool `mapstructure:"complete_other_branches_then_fail"`
	ContinuePipeline              bool `mapstructure:"continue_pipeline"`
	FailOnFailedExpressions       bool `mapstructure:"fail_on_failed_expressions"`
	FailPipeline                  bool `mapstructure:"fail_pipeline"`

	Clusters                          []deployStageCluster   `mapstructure:"cluster"`
	OverrideTimeout                   bool                   `mapstructure:"override_timeout"`
	RestrictExecutionDuringTimeWindow bool                   `mapstructure:"restrict_execution_during_time_window"`
	RestrictedExecutionWindow         []stageExecutionWindow `mapstructure:"restricted_execution_window"`
	StageEnabled                      []stageEnabled         `mapstructure:"stage_enabled"`
}

func newDeployStage() *deployStage {
	return &deployStage{Type: client.DeployStageType}
}

func (s *deployStage) toClientStage() client.Stage {
	cs := client.NewDeployStage()
	cs.Name = s.Name
	cs.RefID = s.RefID
	cs.RequisiteStageRefIds = s.RequisiteStageRefIds
	cs.CompleteOtherBranchesThenFail = s.CompleteOtherBranchesThenFail
	cs.ContinuePipeline = s.ContinuePipeline
	cs.FailOnFailedExpressions = s.FailOnFailedExpressions
	cs.FailPipeline = s.FailPipeline

	for _, c := range s.Clusters {
		cs.Clusters = append(cs.Clusters, *c.toClientCluster())
	}

	cs.OverrideTimeout = s.OverrideTimeout
	cs.RestrictExecutionDuringTimeWindow = s.RestrictExecutionDuringTimeWindow
	if len(s.RestrictedExecutionWindow) > 0 {
		cs.RestrictedExecutionWindow = *s.RestrictedExecutionWindow[0].toClientExecutionWindow()
	}
	if len(s.StageEnabled) > 0 {
		newStageEnabled := client.StageEnabled(s.StageEnabled[0])
		cs.StageEnabled = &newStageEnabled
	}

	return cs
}

// TODO can we just update the ptr?
func (s *deployStage) fromClientStage(cs client.Stage) stage {
	clientStage := cs.(*client.DeployStage)
	newStage := newDeployStage()
	newStage.Name = clientStage.Name
	newStage.RefID = clientStage.RefID
	newStage.RequisiteStageRefIds = clientStage.RequisiteStageRefIds
	newStage.CompleteOtherBranchesThenFail = clientStage.CompleteOtherBranchesThenFail
	newStage.ContinuePipeline = clientStage.ContinuePipeline
	newStage.FailOnFailedExpressions = clientStage.FailOnFailedExpressions
	newStage.FailPipeline = clientStage.FailPipeline

	for _, c := range clientStage.Clusters {
		newStage.Clusters = append(newStage.Clusters, *newClusterFromClientCluster(&c))
	}

	newStage.OverrideTimeout = clientStage.OverrideTimeout
	newStage.RestrictExecutionDuringTimeWindow = clientStage.RestrictExecutionDuringTimeWindow

	newStageExecutionWindow := stageExecutionWindow{}
	newStageExecutionWindow = *newStageExecutionWindow.fromClientWindow(&clientStage.RestrictedExecutionWindow)
	newStage.RestrictedExecutionWindow = append(newStage.RestrictedExecutionWindow, newStageExecutionWindow)
	if clientStage.StageEnabled != nil {
		newStage.StageEnabled = append(newStage.StageEnabled, stageEnabled(*clientStage.StageEnabled))
	}

	return newStage
}

func (s *deployStage) SetResourceData(d *schema.ResourceData) {
	// TODO
	d.Set("name", s.Name)
}

func (s *deployStage) SetRefID(id string) {
	s.RefID = id
}

func (s *deployStage) GetRefID() string {
	return s.RefID
}
