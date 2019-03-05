package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

type rollbackClusterStage struct {
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

	TargetHealthyDeployPercentage   int `mapstructure:"target_healthy_deploy_percentage"`
	TargetHealthyRollbackPercentage int `mapstructure:"target_healthy_rollback_percentage"`
}

func newRollbackClusterStage() *rollbackClusterStage {
	return &rollbackClusterStage{
		Type:         client.RollbackClusterStageType,
		FailPipeline: true,
	}
}

func (s *rollbackClusterStage) toClientStage(config *client.Config) (client.Stage, error) {
	// baseStage
	notifications, err := toClientNotifications(s.Notifications)
	if err != nil {
		return nil, err
	}

	cs := client.NewRollbackClusterStage()
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
	cs.TargetHealthyDeployPercentage = s.TargetHealthyDeployPercentage
	cs.TargetHealthyRollbackPercentage = s.TargetHealthyRollbackPercentage

	return cs, nil
}

func (s *rollbackClusterStage) fromClientStage(cs client.Stage) stage {
	clientStage := cs.(*client.RollbackClusterStage)
	newStage := newRollbackClusterStage()

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
	newStage.TargetHealthyDeployPercentage = clientStage.TargetHealthyDeployPercentage
	newStage.TargetHealthyRollbackPercentage = clientStage.TargetHealthyRollbackPercentage

	return newStage
}

func (s *rollbackClusterStage) SetResourceData(d *schema.ResourceData) error {
	// baseStage
	err := d.Set("name", s.Name)
	if err != nil {
		return nil
	}
	d.Set("requisite_stage_ref_ids", s.RequisiteStageRefIds)
	if err != nil {
		return nil
	}
	d.Set("notification", s.Notifications)
	if err != nil {
		return nil
	}
	d.Set("stage_enabled", s.StageEnabled)
	if err != nil {
		return nil
	}
	d.Set("complete_other_branches_then_fail", s.CompleteOtherBranchesThenFail)
	if err != nil {
		return nil
	}
	d.Set("continue_pipeline", s.ContinuePipeline)
	if err != nil {
		return nil
	}
	err = d.Set("fail_on_failed_expressions", s.FailOnFailedExpressions)
	if err != nil {
		return err
	}
	err = d.Set("fail_pipeline", s.FailPipeline)
	if err != nil {
		return err
	}
	err = d.Set("override_timeout", s.OverrideTimeout)
	if err != nil {
		return err
	}
	err = d.Set("restrict_execution_during_time_window", s.RestrictExecutionDuringTimeWindow)
	if err != nil {
		return err
	}
	err = d.Set("restricted_execution_window", s.RestrictedExecutionWindow)
	if err != nil {
		return err
	}
	// End baseStage

	err = d.Set("cloud_provider", s.CloudProvider)
	if err != nil {
		return err
	}
	err = d.Set("cloud_provider_type", s.CloudProviderType)
	if err != nil {
		return err
	}
	err = d.Set("cluster", s.Cluster)
	if err != nil {
		return err
	}
	err = d.Set("credentials", s.Credentials)
	if err != nil {
		return err
	}
	err = d.Set("moniker", s.Moniker)
	if err != nil {
		return err
	}
	err = d.Set("regions", s.Regions)
	if err != nil {
		return err
	}
	err = d.Set("target_healthy_deploy_percentage", s.TargetHealthyDeployPercentage)
	if err != nil {
		return err
	}
	return d.Set("target_healthy_rollback_percentage", s.TargetHealthyRollbackPercentage)
}

func (s *rollbackClusterStage) SetRefID(id string) {
	s.RefID = id
}

func (s *rollbackClusterStage) GetRefID() string {
	return s.RefID
}
