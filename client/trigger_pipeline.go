package client

import "github.com/mitchellh/mapstructure"

// PipelineTriggerType pipeline trigger
var PipelineTriggerType TriggerType = "pipeline"

func init() {
	triggerFactories[PipelineTriggerType] = parsePipelineTrigger
}

// Pipeline Trigger for Pipeline
type PipelineTrigger struct {
	ID        string      `json:"id"`
	Enabled   bool        `json:"enabled"`
	Type      TriggerType `json:"type"`
	RunAsUser string      `json:"runAsUser,omitempty"`

	Application string   `json:"application"`
	Pipeline    string   `json:"pipeline"`
	Status      []string `json:"status"`
}

func NewPipelineTrigger() *PipelineTrigger {
	return &PipelineTrigger{
		Type: PipelineTriggerType,
	}
}

// GetType for Trigger interface
func (t *PipelineTrigger) GetType() TriggerType {
	return t.Type
}

// GetID for Trigger interface
func (t *PipelineTrigger) GetID() string {
	return t.ID
}

func parsePipelineTrigger(triggerMap map[string]interface{}) (Trigger, error) {
	trigger := NewPipelineTrigger()
	if err := mapstructure.Decode(triggerMap, trigger); err != nil {
		return nil, err
	}
	return trigger, nil
}
