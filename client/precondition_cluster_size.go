package client

import (
	"github.com/mitchellh/mapstructure"
)

// PreconditionClusterSizeType type
var PreconditionClusterSizeType PreconditionType = "clusterSize"

func init() {
	preconditionFactory[PreconditionClusterSizeType] = parsePreconditionClusterSize
}

// PreconditionClusterSizeContext context
type PreconditionClusterSizeContext struct {
	Credentials string   `json:"credentials"`
	Expected    int      `json:"expected"`
	Regions     []string `json:"regions"`
}

// PreconditionClusterSize size
type PreconditionClusterSize struct {
	BasePrecondition `mapstructure:",squash"`

	Context PreconditionClusterSizeContext `json:"context"`
}

// NewPreconditionClusterSize new size
func NewPreconditionClusterSize() *PreconditionClusterSize {
	return &PreconditionClusterSize{
		BasePrecondition: *NewBasePrecondition(PreconditionClusterSizeType),
	}
}

func parsePreconditionClusterSize(preconditionMap map[string]interface{}) (Precondition, error) {
	precondition := NewPreconditionClusterSize()
	if err := mapstructure.Decode(preconditionMap, precondition); err != nil {
		return nil, err
	}
	return precondition, nil
}
