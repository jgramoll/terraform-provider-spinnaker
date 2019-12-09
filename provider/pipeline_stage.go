package provider

import (
	"errors"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

type stage interface {
	fromClientStage(client.Stage) stage
	toClientStage(config *client.Config, refID string) (client.Stage, error)
	SetResourceData(*schema.ResourceData) error
}

type baseStage struct {
	Name                              string                   `mapstructure:"name"`
	RequisiteStageRefIds              []string                 `mapstructure:"requisite_stage_ref_ids"`
	Notifications                     *[]*notification         `mapstructure:"notification"`
	StageEnabled                      *[]*stageEnabled         `mapstructure:"stage_enabled"`
	CompleteOtherBranchesThenFail     bool                     `mapstructure:"complete_other_branches_then_fail"`
	ContinuePipeline                  bool                     `mapstructure:"continue_pipeline"`
	FailOnFailedExpressions           bool                     `mapstructure:"fail_on_failed_expressions"`
	FailPipeline                      bool                     `mapstructure:"fail_pipeline"`
	OverrideTimeout                   bool                     `mapstructure:"override_timeout"`
	StageTimeoutMS                    int                      `mapstructure:"stage_timeout_ms"`
	RestrictExecutionDuringTimeWindow bool                     `mapstructure:"restrict_execution_during_time_window"`
	RestrictedExecutionWindow         *[]*stageExecutionWindow `mapstructure:"restricted_execution_window"`
}

func newBaseStage() *baseStage {
	return &baseStage{
		FailPipeline: true,
	}
}

func (s *baseStage) baseToClientStage(cs *client.BaseStage, refID string) error {
	if refID == "" {
		return errors.New("Ref Id cannot be empty")
	}
	notifications, err := toClientNotifications(s.Notifications)
	if err != nil {
		return err
	}

	cs.Name = s.Name
	cs.RefID = refID
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
	return nil
}

func (s *baseStage) baseFromClientStage(clientStage *client.BaseStage) {
	s.Name = clientStage.Name
	s.RequisiteStageRefIds = clientStage.RequisiteStageRefIds
	s.Notifications = fromClientNotifications(clientStage.Notifications)
	s.StageEnabled = fromClientStageEnabled(clientStage.StageEnabled)
	s.CompleteOtherBranchesThenFail = clientStage.CompleteOtherBranchesThenFail
	s.ContinuePipeline = clientStage.ContinuePipeline
	s.FailOnFailedExpressions = clientStage.FailOnFailedExpressions
	s.FailPipeline = clientStage.FailPipeline
	s.OverrideTimeout = clientStage.OverrideTimeout
	s.RestrictExecutionDuringTimeWindow = clientStage.RestrictExecutionDuringTimeWindow
	s.RestrictedExecutionWindow = fromClientExecutionWindow(clientStage.RestrictedExecutionWindow)
}

func (s *baseStage) baseSetResourceData(d *schema.ResourceData) error {
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
	return nil
}
