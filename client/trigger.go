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
	RunAsUser    string `json:"runAsUser,omitempty"`
	Type         string `json:"type"`
}

// GetTrigger by ID
func (p *Pipeline) GetTrigger(triggerID string) (*Trigger, error) {
	for _, trigger := range p.Triggers {
		if trigger.ID == triggerID {
			return trigger, nil
		}
	}
	return nil, ErrTriggerNotFound
}

// AppendTrigger append trigger
func (pipeline *Pipeline) AppendTrigger(trigger *Trigger) {
	pipeline.Triggers = append(pipeline.Triggers, trigger)
}

// UpdateTrigger in pipeline
func (p *Pipeline) UpdateTrigger(trigger *Trigger) error {
	for i, t := range p.Triggers {
		if t.ID == trigger.ID {
			p.Triggers[i] = trigger
			return nil
		}
	}
	return ErrTriggerNotFound
}

// DeleteTrigger in pipeline
func (p *Pipeline) DeleteTrigger(triggerID string) error {
	for i, t := range p.Triggers {
		if t.ID == triggerID {
			p.Triggers = append(p.Triggers[:i], p.Triggers[i+1:]...)
			return nil
		}
	}
	return ErrTriggerNotFound
}
