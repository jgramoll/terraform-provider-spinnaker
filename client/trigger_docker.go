package client

import "github.com/mitchellh/mapstructure"

// DockerTriggerType docker trigger
var DockerTriggerType TriggerType = "docker"

func init() {
	triggerFactories[DockerTriggerType] = parseDockerTrigger
}

// DockerTrigger for Pipeline
type DockerTrigger struct {
	baseTrigger `mapstructure:",squash"`

	Account      string `json:"account"`
	Organization string `json:"organization"`
	Registry     string `json:"registry"`
	Repository   string `json:"repository"`
	Tag          string `json:"tag"`
}

// NewDockerTrigger new trigger
func NewDockerTrigger() *DockerTrigger {
	return &DockerTrigger{
		baseTrigger: *newBaseTrigger(DockerTriggerType),
	}
}

func parseDockerTrigger(triggerMap map[string]interface{}) (Trigger, error) {
	trigger := NewDockerTrigger()
	if err := mapstructure.Decode(triggerMap, trigger); err != nil {
		return nil, err
	}
	return trigger, nil
}
