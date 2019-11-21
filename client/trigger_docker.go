package client

import "github.com/mitchellh/mapstructure"

// DockerTriggerType docker trigger
var DockerTriggerType TriggerType = "docker"

func init() {
	triggerFactories[DockerTriggerType] = parseDockerTrigger
}

// Docker Trigger for Pipeline
type DockerTrigger struct {
	ID        string      `json:"id"`
	Enabled   bool        `json:"enabled"`
	RunAsUser string      `json:"runAsUser,omitempty"`
	Type      TriggerType `json:"type"`

	Account      string `json:"account"`
	Organization string `json:"organization"`
	Repository   string `json:"repository"`
}

func NewDockerTrigger() *DockerTrigger {
	return &DockerTrigger{
		Type: DockerTriggerType,
	}
}

// GetType for Trigger interface
func (t *DockerTrigger) GetType() TriggerType {
	return t.Type
}

// GetID for Trigger interface
func (t *DockerTrigger) GetID() string {
	return t.ID
}

func parseDockerTrigger(triggerMap map[string]interface{}) (Trigger, error) {
	trigger := NewDockerTrigger()
	if err := mapstructure.Decode(triggerMap, trigger); err != nil {
		return nil, err
	}
	return trigger, nil
}
