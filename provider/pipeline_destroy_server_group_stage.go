package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

type destroyServerGroupStage struct {
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

	CloudProvider     string      `mapstructure:"cloud_provider"`
	CloudProviderType string      `mapstructure:"cloud_provider_type"`
	Cluster           string      `mapstructure:"cluster"`
	Credentials       string      `mapstructure:"credentials"`
	Moniker           *[]*moniker `mapstructure:"moniker"`
	Regions           []string    `mapstructure:"regions"`
	Target            string      `mapstructure:"target"`
}

func newDestroyServerGroupStage() *destroyServerGroupStage {
	return &destroyServerGroupStage{
		Type:         client.DestroyServerGroupStageType,
		FailPipeline: true,
	}
}

func (s *destroyServerGroupStage) toClientStage() (client.Stage, error) {
	// baseStage
	notifications, err := toClientNotifications(s.Notifications)
	if err != nil {
		return nil, err
	}

	cs := client.NewDestroyServerGroupStage()
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

	cs.CloudProvider = s.CloudProvider
	cs.CloudProviderType = s.CloudProviderType
	cs.Cluster = s.Cluster
	cs.Credentials = s.Credentials
	cs.Moniker = toClientMoniker(s.Moniker)
	cs.Regions = s.Regions
	cs.Target = s.Target

	return cs, nil
}

func (s *destroyServerGroupStage) fromClientStage(cs client.Stage) stage {
	clientStage := cs.(*client.DestroyServerGroupStage)
	newStage := newDestroyServerGroupStage()

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

	newStage.CloudProvider = clientStage.CloudProvider
	newStage.CloudProviderType = clientStage.CloudProviderType
	newStage.Cluster = clientStage.Cluster
	newStage.Credentials = clientStage.Credentials
	newStage.Moniker = fromClientMoniker(clientStage.Moniker)
	newStage.Regions = clientStage.Regions
	newStage.Target = clientStage.Target

	return newStage
}

func (s *destroyServerGroupStage) SetResourceData(d *schema.ResourceData) {
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

	d.Set("cloud_provider", s.CloudProvider)
	d.Set("cluster", s.Cluster)
	d.Set("credentials", s.Credentials)
	d.Set("moniker", s.Moniker)
	d.Set("regions", s.Regions)
	d.Set("target", s.Target)
}

func (s *destroyServerGroupStage) SetRefID(id string) {
	s.RefID = id
}

func (s *destroyServerGroupStage) GetRefID() string {
	return s.RefID
}
