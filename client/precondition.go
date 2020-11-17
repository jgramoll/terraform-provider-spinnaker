package client

import (
	"errors"
	"fmt"

	"github.com/mitchellh/mapstructure"
)

// Precondition precondition
type Precondition interface {
	GetType() PreconditionType
}

// BasePrecondition base
type BasePrecondition struct {
	FailPipeline bool             `json:"failPipeline"`
	Type         PreconditionType `json:"type"`
}

// NewBasePrecondition new base
func NewBasePrecondition(t PreconditionType) *BasePrecondition {
	return &BasePrecondition{
		FailPipeline: true,
		Type:         t,
	}
}

// GetType get type
func (p *BasePrecondition) GetType() PreconditionType {
	return p.Type
}

// ParsePreconditions parse
func ParsePreconditions(toParse []interface{}) ([]Precondition, error) {
	preconditions := []Precondition{}

	for _, preconditionMapInterface := range toParse {
		preconditionMap, ok := preconditionMapInterface.(map[string]interface{})
		if !ok {
			return nil, errors.New("invalid precondition")
		}

		typeString, ok := preconditionMap["type"].(string)
		if !ok {
			return nil, errors.New("missing or invalid precondition type")
		}
		preconditionType := PreconditionType(typeString)

		precondition, err := ParsePrecondition(preconditionMap, preconditionType)
		if err != nil {
			return nil, err
		}
		preconditions = append(preconditions, precondition)
	}

	return preconditions, nil
}

// ParsePrecondition parse
func ParsePrecondition(i map[string]interface{}, t PreconditionType) (Precondition, error) {
	preconditionFunc, ok := preconditionFactory[t]
	if !ok {
		return nil, fmt.Errorf("unknown precondition %s", t)
	}

	precondition, err := preconditionFunc(i)
	if err != nil {
		return nil, err
	}
	if err := mapstructure.Decode(i, precondition); err != nil {
		return nil, err
	}
	return precondition, nil
}
