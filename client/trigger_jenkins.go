package client

import "github.com/mitchellh/mapstructure"

// JenkinsTriggerType jenkins trigger
var JenkinsTriggerType TriggerType = "jenkins"

func init() {
	triggerFactories[JenkinsTriggerType] = parseJenkinsTrigger
}

// Jenkins Trigger for Pipeline
type JenkinsTrigger struct {
	baseTrigger `mapstructure:",squash"`

	Job          string `json:"job"`
	Master       string `json:"master"`
	PropertyFile string `json:"propertyFile"`
}

func NewJenkinsTrigger() *JenkinsTrigger {
	return &JenkinsTrigger{
		baseTrigger: *newBaseTrigger(JenkinsTriggerType),
	}
}

func parseJenkinsTrigger(triggerMap map[string]interface{}) (Trigger, error) {
	trigger := NewJenkinsTrigger()
	if err := mapstructure.Decode(triggerMap, trigger); err != nil {
		return nil, err
	}
	return trigger, nil
}
