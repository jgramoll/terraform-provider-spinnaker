package client

import "github.com/mitchellh/mapstructure"

// JenkinsTriggerType jenkins trigger
var JenkinsTriggerType TriggerType = "jenkins"

func init() {
	triggerFactories[JenkinsTriggerType] = parseJenkinsTrigger
}

// Jenkins Trigger for Pipeline
type JenkinsTrigger struct {
	ID           string      `json:"id"`
	Enabled      bool        `json:"enabled"`
	Job          string      `json:"job"`
	Master       string      `json:"master"`
	PropertyFile string      `json:"propertyFile"`
	RunAsUser    string      `json:"runAsUser,omitempty"`
	Type         TriggerType `json:"type"`
}

func NewJenkinsTrigger() *JenkinsTrigger {
	return &JenkinsTrigger{
		Type: JenkinsTriggerType,
	}
}

// GetType for Trigger interface
func (t *JenkinsTrigger) GetType() TriggerType {
	return t.Type
}

// GetID for Trigger interface
func (t *JenkinsTrigger) GetID() string {
	return t.ID
}

func parseJenkinsTrigger(triggerMap map[string]interface{}) (Trigger, error) {
	trigger := NewJenkinsTrigger()
	if err := mapstructure.Decode(triggerMap, trigger); err != nil {
		return nil, err
	}
	return trigger, nil
}
