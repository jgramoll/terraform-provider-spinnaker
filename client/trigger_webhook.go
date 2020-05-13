package client

import "github.com/mitchellh/mapstructure"

// WebhookTriggerType webhook trigger
var WebhookTriggerType TriggerType = "webhook"

func init() {
	triggerFactories[WebhookTriggerType] = parseWebhookTrigger
}

// WebhookTrigger for Pipeline
type WebhookTrigger struct {
	baseTrigger `mapstructure:",squash"`

	Source             string            `json:"source"`
	PayloadConstraints map[string]string `json:"payloadConstraints"`
}

// NewWebhookTrigger new WebhookTrigger
func NewWebhookTrigger() *WebhookTrigger {
	return &WebhookTrigger{
		baseTrigger: *newBaseTrigger(WebhookTriggerType),
	}
}

func parseWebhookTrigger(triggerMap map[string]interface{}) (Trigger, error) {
	trigger := NewWebhookTrigger()
	if err := mapstructure.Decode(triggerMap, trigger); err != nil {
		return nil, err
	}
	return trigger, nil
}
