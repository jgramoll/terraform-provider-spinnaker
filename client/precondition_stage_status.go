package client

import "github.com/mitchellh/mapstructure"

// PreconditionStageStatusType
var PreconditionStageStatusType PreconditionType = "stageStatus"

func init() {
	preconditionFactory[PreconditionStageStatusType] = parsePreconditionStageStatus
}

type PreconditionStageStatusContext struct {
	StageName   string `json:"stageName"`
	StageStatus string `json:"stageStatus"`
}

type PreconditionStageStatus struct {
	BasePrecondition `mapstructure:",squash"`

	Context PreconditionStageStatusContext `json:"context"`
}

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
