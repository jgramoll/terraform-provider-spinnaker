package client

type baseTrigger struct {
	ID        string      `json:"id"`
	Enabled   bool        `json:"enabled"`
	RunAsUser string      `json:"runAsUser,omitempty"`
	Type      TriggerType `json:"type"`
}

func newBaseTrigger(t TriggerType) *baseTrigger {
	return &baseTrigger{
		Type: t,
	}
}

func (t *baseTrigger) GetID() string {
	return t.ID
}

func (t *baseTrigger) GetType() TriggerType {
	return t.Type
}
