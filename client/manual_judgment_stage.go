package client

import (
	"github.com/mitchellh/mapstructure"
)

// ManualJudgmentStageType deploy stage
var ManualJudgmentStageType StageType = "manualJudgment"

func init() {
	stageFactories[ManualJudgmentStageType] = parseManualJudgmentStage
}

// JudgmentInputs inputs for judgment
type JudgmentInputs struct {
	Value string `json:"value"`
}

type ManualJudgmentStage struct {
	BaseStage `mapstructure:",squash"`

	Instructions   string           `json:"instructions"`
	JudgmentInputs []JudgmentInputs `json:"judgmentInputs"`
}

// NewManualJudgmentStage for pipeline
func NewManualJudgmentStage() *ManualJudgmentStage {
	return &ManualJudgmentStage{
		BaseStage: *newBaseStage(ManualJudgmentStageType),
	}
}

func parseManualJudgmentStage(stageMap map[string]interface{}) (Stage, error) {
	stage := NewManualJudgmentStage()
	if err := stage.parseBaseStage(stageMap); err != nil {
		return nil, err
	}

	if err := mapstructure.Decode(stageMap, stage); err != nil {
		return nil, err
	}
	return stage, nil
}
