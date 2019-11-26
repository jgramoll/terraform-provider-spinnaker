package client

import (
	"github.com/mitchellh/mapstructure"
)

// WebhookStageType bake stage
var WebhookStageType StageType = "webhook"

func init() {
	stageFactories[WebhookStageType] = parseWebhookStage
}

// WebhookStage for pipeline
type WebhookStage struct {
	BaseStage `mapstructure:",squash"`

	CanceledStatuses    string                 `json:"canceledStatuses"`
	CustomHeaders       map[string]string      `json:"customHeaders"`
	FailFastStatusCodes []string               `json:"failFastStatusCodes"`
	Method              string                 `json:"method"`
	Payload             map[string]interface{} `json:"payload"`
	ProgressJSONPath    string                 `json:"progressJsonPath"`
	StatusJSONPath      string                 `json:"statusJsonPath"`
	StatusURLJSONPath   string                 `json:"statusUrlJsonPath"`
	StatusURLResolution string                 `json:"statusUrlResolution"`
	SuccessStatuses     string                 `json:"successStatuses"`
	TerminalStatuses    string                 `json:"terminalStatuses"`
	URL                 string                 `json:"url"`
	WaitForCompletion   bool                   `json:"waitForCompletion"`
}

// NewWebhookStage for pipeline
func NewWebhookStage() *WebhookStage {
	return &WebhookStage{
		BaseStage: *newBaseStage(WebhookStageType),
	}
}

func parseWebhookStage(stageMap map[string]interface{}) (Stage, error) {
	stage := NewWebhookStage()
	if err := stage.parseBaseStage(stageMap); err != nil {
		return nil, err
	}

	if err := mapstructure.Decode(stageMap, stage); err != nil {
		return nil, err
	}
	return stage, nil
}
