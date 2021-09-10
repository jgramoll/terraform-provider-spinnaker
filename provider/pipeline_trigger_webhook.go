package provider

import (
	"errors"

	"github.com/get-bridge/terraform-provider-spinnaker/client"
	"github.com/hashicorp/terraform/helper/schema"
)

// Webhook trigger for Pipeline
type webhookTrigger struct {
	baseTrigger `mapstructure:",squash"`

	Source             string            `mapstructure:"source"`
	PayloadConstraints map[string]string `mapstructure:"payload_constraints"`
}

func newWebhookTrigger() *webhookTrigger {
	return &webhookTrigger{}
}

func (t *webhookTrigger) toClientTrigger(id string) (client.Trigger, error) {
	clientTrigger := client.NewWebhookTrigger()
	clientTrigger.ID = id
	clientTrigger.Enabled = t.Enabled

	clientTrigger.Source = t.Source
	clientTrigger.PayloadConstraints = t.PayloadConstraints
	return clientTrigger, nil
}

func (*webhookTrigger) fromClientTrigger(clientTriggerInterface client.Trigger) (trigger, error) {
	clientTrigger, ok := clientTriggerInterface.(*client.WebhookTrigger)
	if !ok {
		return nil, errors.New("Expected pipeline trigger")
	}
	t := newWebhookTrigger()
	t.Enabled = clientTrigger.Enabled

	t.Source = clientTrigger.Source
	t.PayloadConstraints = clientTrigger.PayloadConstraints
	return t, nil
}

func (t *webhookTrigger) setResourceData(d *schema.ResourceData) error {
	var err error
	err = d.Set("enabled", t.Enabled)
	if err != nil {
		return err
	}
	err = d.Set("source", t.Source)
	if err != nil {
		return err
	}
	err = d.Set("payload_constraints", t.PayloadConstraints)
	if err != nil {
		return err
	}
	return nil
}
