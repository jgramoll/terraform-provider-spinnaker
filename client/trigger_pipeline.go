package client

import "github.com/mitchellh/mapstructure"

// PipelineTriggerType pipeline trigger
var PipelineTriggerType TriggerType = "pipeline"

func init() {
	triggerFactories[PipelineTriggerType] = parsePipelineTrigger
}

// PipelineTrigger for Pipeline
type PipelineTrigger struct {
	baseTrigger `mapstructure:",squash"`

	Application string   `json:"application"`
	Pipeline    string   `json:"pipeline"`
	Status      []string `json:"status"`
}

// NewPipelineTrigger new trigger
func NewPipelineTrigger() *PipelineTrigger {
	return &PipelineTrigger{
		baseTrigger: *newBaseTrigger(PipelineTriggerType),
	}
}

func parsePipelineTrigger(triggerMap map[string]interface{}) (Trigger, error) {
	trigger := NewPipelineTrigger()
	if err := mapstructure.Decode(triggerMap, trigger); err != nil {
		return nil, err
	}
	return trigger, nil
}
