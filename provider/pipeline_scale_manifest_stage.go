package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

type scaleManifestStage struct {
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

	Account        string                 `mapstructure:"account"`
	Application    string                 `mapstructure:"application"`
	CloudProvider  string                 `mapstructure:"cloud_provider"`
	Cluster        string                 `mapstructure:"cluster"`
	Criteria       string                 `mapstructure:"criteria"`
	IsNew          bool                   `mapstructure:"is_new"`
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
		Type:                    client.ScaleManifestStageType,
		FailPipeline:            true,
		// TODO defaults
	}
}

func (s *scaleManifestStage) toClientStage(config *client.Config) (client.Stage, error) {
	// baseStage
	notifications, err := toClientNotifications(s.Notifications)
	if err != nil {
		return nil, err
	}

	cs := client.NewScaleManifestStage()
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

	cs.Account = s.Account
	cs.Application = s.Application
	cs.CloudProvider = s.CloudProvider
	cs.Cluster = s.Cluster
	cs.Criteria = s.Criteria
	cs.IsNew = s.IsNew
	cs.Kind = s.Kind
	cs.Kinds = s.Kinds
	cs.LabelSelectors = s.LabelSelectors
	cs.Location = s.Location
	cs.ManifestName = s.ManifestName
	cs.Mode = s.Mode
	cs.Replicas = s.Replicas
	return cs, nil
}

func (s *scaleManifestStage) fromClientStage(cs client.Stage) stage {
	clientStage := cs.(*client.ScaleManifestStage)
	newStage := newScaleManifestStage()

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

	newStage.Account = clientStage.Account
	newStage.Application = clientStage.Application
	newStage.CloudProvider = clientStage.CloudProvider
	newStage.Cluster = clientStage.Cluster
	newStage.Criteria = clientStage.Criteria
	newStage.IsNew = clientStage.IsNew
	newStage.Kind = clientStage.Kind
	newStage.Kinds = clientStage.Kinds
	newStage.LabelSelectors = clientStage.LabelSelectors
	newStage.Location = clientStage.Location
	newStage.ManifestName = clientStage.ManifestName
	newStage.Mode = clientStage.Mode
	newStage.Replicas = clientStage.Replicas

	return newStage
}

func (s *scaleManifestStage) SetResourceData(d *schema.ResourceData) error {
	// baseStage
	err := d.Set("name", s.Name)
	if err != nil {
		return err
	}
	err = d.Set("requisite_stage_ref_ids", s.RequisiteStageRefIds)
	if err != nil {
		return err
	}
	err = d.Set("notification", s.Notifications)
	if err != nil {
		return err
	}
	err = d.Set("stage_enabled", s.StageEnabled)
	if err != nil {
		return err
	}
	err = d.Set("complete_other_branches_then_fail", s.CompleteOtherBranchesThenFail)
	if err != nil {
		return err
	}
	err = d.Set("continue_pipeline", s.ContinuePipeline)
	if err != nil {
		return err
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
	err = d.Set("is_new", s.IsNew)
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

func (s *scaleManifestStage) SetRefID(id string) {
	s.RefID = id
}

func (s *scaleManifestStage) GetRefID() string {
	return s.RefID
}
