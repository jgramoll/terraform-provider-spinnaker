package provider

import (
	"errors"

	"github.com/get-bridge/terraform-provider-spinnaker/client"
	"github.com/hashicorp/terraform/helper/schema"
)

// Pipeline trigger for Pipeline
type pipelineTrigger struct {
	baseTrigger `mapstructure:",squash"`

	Application string   `mapstructure:"triggering_application"`
	Pipeline    string   `mapstructure:"triggering_pipeline"`
	Status      []string `mapstructure:"status"`
}

func newPipelineTrigger() *pipelineTrigger {
	return &pipelineTrigger{}
}

func (t *pipelineTrigger) toClientTrigger(id string) (client.Trigger, error) {
	clientTrigger := client.NewPipelineTrigger()
	clientTrigger.ID = id
	clientTrigger.Enabled = t.Enabled

	clientTrigger.Application = t.Application
	clientTrigger.Pipeline = t.Pipeline
	clientTrigger.Status = t.Status
	return clientTrigger, nil
}

func (*pipelineTrigger) fromClientTrigger(clientTriggerInterface client.Trigger) (trigger, error) {
	clientTrigger, ok := clientTriggerInterface.(*client.PipelineTrigger)
	if !ok {
		return nil, errors.New("Expected pipeline trigger")
	}
	t := newPipelineTrigger()
	t.Enabled = clientTrigger.Enabled

	t.Application = clientTrigger.Application
	t.Pipeline = clientTrigger.Pipeline
	t.Status = clientTrigger.Status
	return t, nil
}

func (t *pipelineTrigger) setResourceData(d *schema.ResourceData) error {
	var err error
	err = d.Set("enabled", t.Enabled)
	if err != nil {
		return err
	}
	err = d.Set("triggering_application", t.Application)
	if err != nil {
		return err
	}
	err = d.Set("triggering_pipeline", t.Pipeline)
	if err != nil {
		return err
	}
	err = d.Set("status", t.Status)
	if err != nil {
		return err
	}
	return nil
}
