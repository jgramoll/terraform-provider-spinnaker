package client

import "github.com/mitchellh/mapstructure"

// PreconditionStageStatusType type
var PreconditionStageStatusType PreconditionType = "stageStatus"

func init() {
	preconditionFactory[PreconditionStageStatusType] = parsePreconditionStageStatus
}

// PreconditionStageStatusContext context
type PreconditionStageStatusContext struct {
	StageName   string `json:"stageName"`
	StageStatus string `json:"stageStatus"`
}

// PreconditionStageStatus status
type PreconditionStageStatus struct {
	BasePrecondition `mapstructure:",squash"`

	Context PreconditionStageStatusContext `json:"context"`
}

// NewPreconditionStageStatus new status
func NewPreconditionStageStatus() *PreconditionStageStatus {
	return &PreconditionStageStatus{
		BasePrecondition: *NewBasePrecondition(PreconditionStageStatusType),
	}
}

func parsePreconditionStageStatus(preconditionMap map[string]interface{}) (Precondition, error) {
	precondition := NewPreconditionStageStatus()
	if err := mapstructure.Decode(preconditionMap, precondition); err != nil {
		return nil, err
	}
	return precondition, nil
}
