package client

var triggerFactories = map[TriggerType]func(map[string]interface{}) (Trigger, error){}

// TriggerType type of trigger
type TriggerType string

func (st TriggerType) String() string {
	return string(st)
}
