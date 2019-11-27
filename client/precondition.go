package client

import (
	"errors"
	"fmt"

	"github.com/mitchellh/mapstructure"
)

type Precondition interface {
	GetType() PreconditionType
}

type BasePrecondition struct {
	FailPipeline bool             `json:"failPipeline"`
	Type         PreconditionType `json:"type"`
}

func NewBasePrecondition(t PreconditionType) *BasePrecondition {
	return &BasePrecondition{
		FailPipeline: true,
		Type:         t,
	}
}

func (p *BasePrecondition) GetType() PreconditionType {
	return p.Type
}

func ParsePreconditions(toParse []map[string]interface{}) ([]Precondition, error) {
	preconditions := []Precondition{}

	for _, preconditionMap := range toParse {
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
