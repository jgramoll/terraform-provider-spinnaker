package client

import (
	"github.com/mitchellh/mapstructure"
)

// PreconditionExpressionType type
var PreconditionExpressionType PreconditionType = "expression"

func init() {
	preconditionFactory[PreconditionExpressionType] = parsePreconditionExpression
}

// PreconditionExpressionContext context
type PreconditionExpressionContext struct {
	Expression string `json:"expression"`
}

// PreconditionExpression expression
type PreconditionExpression struct {
	BasePrecondition `mapstructure:",squash"`

	Context PreconditionExpressionContext `json:"context"`
}

// NewPreconditionExpression new expression
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
