package client

import (
	"errors"
)

// ErrTriggerNotFound trigger not found
var ErrTriggerNotFound = errors.New("could not find trigger")

// Trigger for Pipeline
type Trigger struct {
	ID           string `json:"id"`
	Enabled      bool   `json:"enabled"`
	Job          string `json:"job"`
	Master       string `json:"master"`
	PropertyFile string `json:"propertyFile"`
	Type         string `json:"type"`
}

// GetTrigger by ID
func (p *Pipeline) GetTrigger(triggerID string) (*Trigger, error) {
	if p.Triggers != nil {
		for _, trigger := range *p.Triggers {
			if trigger.ID == triggerID {
				return trigger, nil
			}
		}
	}
	return nil, ErrTriggerNotFound
}

// UpdateTrigger in pipeline
func (p *Pipeline) UpdateTrigger(trigger *Trigger) error {
	if p.Triggers != nil {
		for i, t := range *p.Triggers {
			if t.ID == trigger.ID {
				(*p.Triggers)[i] = trigger
				return nil
			}
		}
	}
	return ErrTriggerNotFound
}

// DeleteTrigger in pipeline
func (p *Pipeline) DeleteTrigger(triggerID string) error {
	if p.Triggers != nil {
		triggers := *p.Triggers
		for i, t := range triggers {
			if t.ID == triggerID {
				triggers = append(triggers[:i], triggers[i+1:]...)
				p.Triggers = &triggers
				return nil
			}
		}
	}
	return ErrTriggerNotFound
}
