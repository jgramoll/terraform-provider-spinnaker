package client

import (
	"github.com/mitchellh/mapstructure"
)

// WebhookStageType bake stage
var WebhookStageType StageType = "webhook"

func init() {
	stageFactories[WebhookStageType] = parseWebhookStage
}

type serializableWebhookStage struct {
	// BaseStage
	Name                              string                `json:"name"`
	RefID                             string                `json:"refId"`
	Type                              StageType             `json:"type"`
	RequisiteStageRefIds              []string              `json:"requisiteStageRefIds"`
	SendNotifications                 bool                  `json:"sendNotifications"`
	StageEnabled                      *StageEnabled         `json:"stageEnabled"`
	CompleteOtherBranchesThenFail     bool                  `json:"completeOtherBranchesThenFail"`
	ContinuePipeline                  bool                  `json:"continuePipeline"`
	FailOnFailedExpressions           bool                  `json:"failOnFailedExpressions"`
	FailPipeline                      bool                  `json:"failPipeline"`
	OverrideTimeout                   bool                  `json:"overrideTimeout"`
	RestrictExecutionDuringTimeWindow bool                  `json:"restrictExecutionDuringTimeWindow"`
	RestrictedExecutionWindow         *StageExecutionWindow `json:"restrictedExecutionWindow"`
	// End BaseStage

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
	WaitForCompletion   bool              `mapstructure:"wait_for_completion"`
}

// WebhookStage for pipeline
type WebhookStage struct {
	*serializableWebhookStage
	Notifications *[]*Notification `json:"notifications"`
}

func newserializableWebhookStage() *serializableWebhookStage {
	return &serializableWebhookStage{
		Type:                 WebhookStageType,
		FailPipeline:         true,
		RequisiteStageRefIds: []string{},
	}
}

// NewWebhookStage for pipeline
func NewWebhookStage() *WebhookStage {
	return &WebhookStage{
		serializableWebhookStage: newserializableWebhookStage(),
	}
}

// GetName for Stage interface
func (s *WebhookStage) GetName() string {
	return s.Name
}

// GetType for Stage interface
func (s *WebhookStage) GetType() StageType {
	return s.Type
}

// GetRefID for Stage interface
func (s *WebhookStage) GetRefID() string {
	return s.RefID
}

func parseWebhookStage(stageMap map[string]interface{}) (Stage, error) {
	stage := newserializableWebhookStage()
	if err := mapstructure.Decode(stageMap, stage); err != nil {
		return nil, err
	}

	notifications, err := parseNotifications(stageMap["notifications"])
	if err != nil {
		return nil, err
	}
	return &WebhookStage{
		serializableWebhookStage: stage,
		Notifications:                      notifications,
	}, nil
}
