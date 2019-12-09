package client

import (
	"github.com/mitchellh/mapstructure"
)

// EvaluateVariablesStageType bake stage
var EvaluateVariablesStageType StageType = "evaluateVariables"

func init() {
	stageFactories[EvaluateVariablesStageType] = parseEvaluateVariablesStage
}

// Variable for evaluation
type Variable struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// EvaluateVariablesStage for pipeline
type EvaluateVariablesStage struct {
	BaseStage `mapstructure:",squash"`

	Variables []Variable `json:"variables"`
}

// NewEvaluateVariablesStage for pipeline
func NewEvaluateVariablesStage() *EvaluateVariablesStage {
	return &EvaluateVariablesStage{
		BaseStage: *newBaseStage(EvaluateVariablesStageType),
	}
}

func parseEvaluateVariablesStage(stageMap map[string]interface{}) (Stage, error) {
	stage := NewEvaluateVariablesStage()
	if err := stage.parseBaseStage(stageMap); err != nil {
		return nil, err
	}

	if err := mapstructure.Decode(stageMap, stage); err != nil {
		return nil, err
	}
	return stage, nil
}
