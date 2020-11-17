package client

import (
	"errors"
	"log"
)

// ErrTriggerNotFound trigger not found
var ErrTriggerNotFound = errors.New("could not find trigger")

// Trigger interface for Pipeline triggers
type Trigger interface {
	GetID() string
	GetType() TriggerType
}

// GetTrigger by ID
func (p *Pipeline) GetTrigger(triggerID string) (Trigger, error) {
	for _, trigger := range p.Triggers {
		if trigger.GetID() == triggerID {
			return trigger, nil
		}
	}
	return nil, ErrTriggerNotFound
}

// AppendTrigger append trigger
func (p *Pipeline) AppendTrigger(trigger Trigger) {
	p.Triggers = append(p.Triggers, trigger)
}

// UpdateTrigger in pipeline
func (p *Pipeline) UpdateTrigger(trigger Trigger) error {
	for i, t := range p.Triggers {
		if t.GetID() == trigger.GetID() {
			p.Triggers[i] = trigger
			return nil
		}
	}
	return ErrTriggerNotFound
}

// DeleteTrigger in pipeline
func (p *Pipeline) DeleteTrigger(triggerID string) error {
	for i, t := range p.Triggers {
		if t.GetID() == triggerID {
			p.Triggers = append(p.Triggers[:i], p.Triggers[i+1:]...)
			return nil
		}
	}
	return ErrTriggerNotFound
}

func parseTriggers(triggersHashInterface interface{}) (*[]Trigger, error) {
	triggers := []Trigger{}
	if triggersHashInterface == nil {
		return &triggers, nil
	}

	triggersToParse := triggersHashInterface.([]interface{})
	for _, triggerInterface := range triggersToParse {
		triggerMap := triggerInterface.(map[string]interface{})

		triggerTypeInterface, ok := triggerMap["type"]
		if !ok {
			log.Println("[WARN] pipeline trigger type is missing")
			continue
		}
		triggerType := TriggerType(triggerTypeInterface.(string))

		factory := triggerFactories[triggerType]
		if factory == nil {
			log.Printf("[WARN] unknown pipeline trigger \"%s\"\n", triggerType)
			continue
		}
		trigger, err := factory(triggerMap)
		if err != nil {
			return nil, err
		}
		triggers = append(triggers, trigger)
	}
	return &triggers, nil
}
