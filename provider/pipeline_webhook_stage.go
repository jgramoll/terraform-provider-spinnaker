package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

type webhookStage struct {
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

	CanceledStatuses    string            `mapstructure:"canceled_statuses"`
	CustomHeaders       map[string]string `mapstructure:"custom_headers"`
	FailFastStatusCodes []string          `mapstructure:"fail_fast_status_codes"`
	Method              string            `mapstructure:"method"`
	Payload             string            `mapstructure:"payload"`
	ProgressJSONPath    string            `mapstructure:"progress_json_path"`
	StatusJSONPath      string            `mapstructure:"status_json_path"`
	StatusURLJSONPath   string            `mapstructure:"status_url_json_path"`
	StatusURLResolution string            `mapstructure:"status_url_resolution"`
	SuccessStatuses     string            `mapstructure:"success_statuses"`
	TerminalStatuses    string            `mapstructure:"terminal_statuses"`
	URL                 string            `mapstructure:"url"`
}

func newWebhookStage() *webhookStage {
	return &webhookStage{
		Type:         client.WebhookStageType,
		FailPipeline: true,
	}
}

func (s *webhookStage) toClientStage(config *client.Config) (client.Stage, error) {
	// baseStage
	notifications, err := toClientNotifications(s.Notifications)
	if err != nil {
		return nil, err
	}

	cs := client.NewWebhookStage()
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

	cs.CanceledStatuses = s.CanceledStatuses
	cs.CustomHeaders = s.CustomHeaders
	cs.FailFastStatusCodes = s.FailFastStatusCodes
	cs.Method = s.Method
	cs.Payload = s.Payload
	cs.ProgressJSONPath = s.ProgressJSONPath
	cs.StatusJSONPath = s.StatusJSONPath
	cs.StatusURLJSONPath = s.StatusURLJSONPath
	cs.StatusURLResolution = s.StatusURLResolution
	cs.SuccessStatuses = s.SuccessStatuses
	cs.TerminalStatuses = s.TerminalStatuses
	cs.URL = s.URL
	cs.WaitForCompletion = s.StatusURLResolution != ""

	return cs, nil
}

func (s *webhookStage) fromClientStage(cs client.Stage) stage {
	clientStage := cs.(*client.WebhookStage)
	newStage := newWebhookStage()

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

	newStage.CanceledStatuses = clientStage.CanceledStatuses
	newStage.CustomHeaders = clientStage.CustomHeaders
	newStage.FailFastStatusCodes = clientStage.FailFastStatusCodes
	newStage.Method = clientStage.Method
	newStage.Payload = clientStage.Payload
	newStage.ProgressJSONPath = clientStage.ProgressJSONPath
	newStage.StatusJSONPath = clientStage.StatusJSONPath
	newStage.StatusURLJSONPath = clientStage.StatusURLJSONPath
	newStage.StatusURLResolution = clientStage.StatusURLResolution
	newStage.SuccessStatuses = clientStage.SuccessStatuses
	newStage.TerminalStatuses = clientStage.TerminalStatuses
	newStage.URL = clientStage.URL

	return newStage
}

func (s *webhookStage) SetResourceData(d *schema.ResourceData) error {
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

	err = d.Set("canceled_statuses", s.CanceledStatuses)
	if err != nil {
		return err
	}
	err = d.Set("custom_headers", s.CustomHeaders)
	if err != nil {
		return err
	}
	err = d.Set("fail_fast_status_codes", s.FailFastStatusCodes)
	if err != nil {
		return err
	}
	err = d.Set("method", s.Method)
	if err != nil {
		return err
	}
	err = d.Set("payload", s.Payload)
	if err != nil {
		return err
	}
	err = d.Set("progress_json_path", s.ProgressJSONPath)
	if err != nil {
		return err
	}
	err = d.Set("status_json_path", s.StatusJSONPath)
	if err != nil {
		return err
	}
	err = d.Set("status_url_json_path", s.StatusURLJSONPath)
	if err != nil {
		return err
	}
	err = d.Set("status_url_resolution", s.StatusURLResolution)
	if err != nil {
		return err
	}
	err = d.Set("success_statuses", s.SuccessStatuses)
	if err != nil {
		return err
	}
	err = d.Set("terminal_statuses", s.TerminalStatuses)
	if err != nil {
		return err
	}
	return d.Set("url", s.URL)

}

func (s *webhookStage) SetRefID(id string) {
	s.RefID = id
}

func (s *webhookStage) GetRefID() string {
	return s.RefID
}
