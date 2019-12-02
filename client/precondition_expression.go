package client

import (
	"github.com/mitchellh/mapstructure"
)

// PreconditionExpressionType
var PreconditionExpressionType PreconditionType = "expression"

func init() {
	preconditionFactory[PreconditionExpressionType] = parsePreconditionExpression
}

type PreconditionExpressionContext struct {
	Expression string `json:"expression"`
}

type PreconditionExpression struct {
	BasePrecondition `mapstructure:",squash"`

	Context PreconditionExpressionContext `json:"context"`
}

func NewPreconditionExpression() *PreconditionExpression {
	return &PreconditionExpression{
		BasePrecondition: *NewBasePrecondition(PreconditionExpressionType),
	}
}

func parsePreconditionExpression(preconditionMap map[string]interface{}) (Precondition, error) {
	precondition := NewPreconditionExpression()

	if err := mapstructure.Decode(preconditionMap, precondition); err != nil {
		return nil, err
	}
	return precondition, nil
}
