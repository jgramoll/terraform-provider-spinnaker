package provider

import (
	"encoding/json"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

type webhookStage struct {
	baseStage `mapstructure:",squash"`

	CanceledStatuses    string            `mapstructure:"canceled_statuses"`
	CustomHeaders       map[string]string `mapstructure:"custom_headers"`
	FailFastStatusCodes []string          `mapstructure:"fail_fast_status_codes"`
	Method              string            `mapstructure:"method"`
	PayloadString       string            `mapstructure:"payload_string"`
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
		baseStage: *newBaseStage(),
	}
}

func (s *webhookStage) toClientStage(config *client.Config, refID string) (client.Stage, error) {
	cs := client.NewWebhookStage()
	err := s.baseToClientStage(&cs.BaseStage, refID, newDefaultNotificationInterface)
	if err != nil {
		return nil, err
	}

	cs.CanceledStatuses = s.CanceledStatuses
	cs.CustomHeaders = s.CustomHeaders
	cs.FailFastStatusCodes = s.FailFastStatusCodes
	cs.Method = s.Method

	var definedPayload map[string]interface{}
	if s.PayloadString != "" {
		if err = json.Unmarshal([]byte(s.PayloadString), &definedPayload); err != nil {
			return nil, err
		}
	}
	cs.Payload = definedPayload

	cs.ProgressJSONPath = s.ProgressJSONPath
	cs.StatusJSONPath = s.StatusJSONPath
	cs.StatusURLJSONPath = s.StatusURLJSONPath
	statusResolution := s.StatusURLResolution
	if s.StatusURLResolution == "" {
		statusResolution = "getMethod"
	}
	cs.StatusURLResolution = statusResolution
	cs.SuccessStatuses = s.SuccessStatuses
	cs.TerminalStatuses = s.TerminalStatuses
	cs.URL = s.URL
	cs.WaitForCompletion = s.StatusURLResolution != ""

	return cs, nil
}

func (*webhookStage) fromClientStage(cs client.Stage) (stage, error) {
	clientStage := cs.(*client.WebhookStage)
	newStage := newWebhookStage()
	err := newStage.baseFromClientStage(&clientStage.BaseStage, newDefaultNotificationInterface)
	if err != nil {
		return nil, err
	}

	newStage.CanceledStatuses = clientStage.CanceledStatuses
	newStage.CustomHeaders = clientStage.CustomHeaders
	newStage.FailFastStatusCodes = clientStage.FailFastStatusCodes
	newStage.Method = clientStage.Method

	if clientStage.Payload != nil && len(clientStage.Payload) > 0 {
		out, err := json.Marshal(clientStage.Payload)
		if err != nil {
			log.Println("[WARN]: Failed to unmarshal payload into string")
		}
		newStage.PayloadString = string(out)
	}

	newStage.ProgressJSONPath = clientStage.ProgressJSONPath
	newStage.StatusJSONPath = clientStage.StatusJSONPath
	newStage.StatusURLJSONPath = clientStage.StatusURLJSONPath
	statusResolution := clientStage.StatusURLResolution
	if !clientStage.WaitForCompletion {
		statusResolution = ""
	}
	newStage.StatusURLResolution = statusResolution
	newStage.SuccessStatuses = clientStage.SuccessStatuses
	newStage.TerminalStatuses = clientStage.TerminalStatuses
	newStage.URL = clientStage.URL

	return newStage, nil
}

func (s *webhookStage) SetResourceData(d *schema.ResourceData) error {
	err := s.baseSetResourceData(d)
	if err != nil {
		return err
	}

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
	err = d.Set("payload_string", s.PayloadString)
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
