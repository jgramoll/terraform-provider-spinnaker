package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

type deployStage struct {
	// baseStage
	Name                              string                   `mapstructure:"name"`
	RefID                             string                   `mapstructure:"ref_id"`
	Type                              client.StageType         `mapstructure:"type"`
	RequisiteStageRefIds              []string                 `mapstructure:"requisite_stage_ref_ids"`
	Notifications                     *[]*notification         `mapstructure:"notification"`
	StageEnabled                      *[]*stageEnabled         `mapstructure:"stage_enabled"`
	CompleteOtherBranchesThenFail     bool                     `mapstructure:"complete_other_branches_then_fail"`
	ContinuePipeline                  bool                     `mapstructure:"continue_pipeline"`
	FailOnFailedExpressions           bool                     `mapstructure:"fail_on_failed_expressions"`
	FailPipeline                      bool                     `mapstructure:"fail_pipeline"`
	OverrideTimeout                   bool                     `mapstructure:"override_timeout"`
	RestrictExecutionDuringTimeWindow bool                     `mapstructure:"restrict_execution_during_time_window"`
	RestrictedExecutionWindow         *[]*stageExecutionWindow `mapstructure:"restricted_execution_window"`
	// End baseStage

	Clusters *[]*deployStageCluster `mapstructure:"cluster"`
}

func newDeployStage() *deployStage {
	return &deployStage{
		Type:         client.DeployStageType,
		FailPipeline: true,
	}
}

func (s *deployStage) toClientStage() (client.Stage, error) {
	// baseStage
	notifications, err := toClientNotifications(s.Notifications)
	if err != nil {
		return nil, err
	}

	cs := client.NewDeployStage()
	cs.Name = s.Name
	cs.RefID = s.RefID
	cs.RequisiteStageRefIds = s.RequisiteStageRefIds
	cs.Notifications = notifications
	cs.SendNotifications = notifications != nil && len(*notifications) > 0
	cs.StageEnabled = toClientStageEnabled(s.StageEnabled)
	cs.CompleteOtherBranchesThenFail = s.CompleteOtherBranchesThenFail
	cs.ContinuePipeline = s.ContinuePipeline
	cs.FailOnFailedExpressions = s.FailOnFailedExpressions
	cs.FailPipeline = s.FailPipeline
	cs.OverrideTimeout = s.OverrideTimeout
	cs.RestrictExecutionDuringTimeWindow = s.RestrictExecutionDuringTimeWindow
	cs.RestrictedExecutionWindow = toClientExecutionWindow(s.RestrictedExecutionWindow)
	// End baseStage

	cs.Clusters = s.toClientClusters()

	return cs, nil
}

func (s *deployStage) fromClientStage(cs client.Stage) stage {
	clientStage := cs.(*client.DeployStage)
	newStage := newDeployStage()

	// baseStage
	newStage.Name = clientStage.Name
	newStage.RefID = clientStage.RefID
	newStage.RequisiteStageRefIds = clientStage.RequisiteStageRefIds
	newStage.Notifications = fromClientNotifications(clientStage.Notifications)
	newStage.StageEnabled = fromClientStageEnabled(clientStage.StageEnabled)
	newStage.CompleteOtherBranchesThenFail = clientStage.CompleteOtherBranchesThenFail
	newStage.ContinuePipeline = clientStage.ContinuePipeline
	newStage.FailOnFailedExpressions = clientStage.FailOnFailedExpressions
	newStage.FailPipeline = clientStage.FailPipeline
	newStage.OverrideTimeout = clientStage.OverrideTimeout
	newStage.RestrictExecutionDuringTimeWindow = clientStage.RestrictExecutionDuringTimeWindow
	newStage.RestrictedExecutionWindow = fromClientExecutionWindow(clientStage.RestrictedExecutionWindow)
	// end baseStage

	newStage.Clusters = fromClientClusters(clientStage.Clusters)

	return newStage
}

func (s *deployStage) SetResourceData(d *schema.ResourceData) {
	// baseStage
	d.Set("name", s.Name)
	d.Set("ref_id", s.RefID)
	d.Set("requisite_stage_ref_ids", s.RequisiteStageRefIds)
	d.Set("notification", s.Notifications)
	d.Set("stage_enabled", s.StageEnabled)
	d.Set("complete_other_branches_then_fail", s.CompleteOtherBranchesThenFail)
	d.Set("continue_pipeline", s.ContinuePipeline)
	d.Set("fail_on_failed_expressions", s.FailOnFailedExpressions)
	d.Set("fail_pipeline", s.FailPipeline)
	d.Set("override_timeout", s.OverrideTimeout)
	d.Set("restrict_execution_during_time_window", s.RestrictExecutionDuringTimeWindow)
	d.Set("restricted_execution_window", s.RestrictedExecutionWindow)
	// End baseStage

	d.Set("cluster", s.Clusters)
}

func (s *deployStage) SetRefID(id string) {
	s.RefID = id
}

func (s *deployStage) GetRefID() string {
	return s.RefID
}
