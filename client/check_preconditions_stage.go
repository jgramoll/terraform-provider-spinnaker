package client

import (
	"fmt"
	"reflect"

	"github.com/mitchellh/mapstructure"
)

// CheckPreconditionsStageType check preconditions stage
var CheckPreconditionsStageType StageType = "checkPreconditions"

func init() {
	stageFactories[CheckPreconditionsStageType] = parseCheckPreconditionsStage
}

// CheckPreconditionsStage preconditions stage
type CheckPreconditionsStage struct {
	BaseStage `mapstructure:",squash"`

	Preconditions []Precondition `json:"preconditions"`
}

// NewCheckPreconditionsStage new stage
func NewCheckPreconditionsStage() *CheckPreconditionsStage {
	return &CheckPreconditionsStage{
		BaseStage:     *newBaseStage(CheckPreconditionsStageType),
		Preconditions: []Precondition{},
	}
}

func parseCheckPreconditionsStage(stageMap map[string]interface{}) (Stage, error) {
	stage := NewCheckPreconditionsStage()
	if err := stage.parseBaseStage(stageMap); err != nil {
		return nil, err
	}

	preconditionsInterface, ok := stageMap["preconditions"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("Could not parse preconditions: %v", reflect.TypeOf(stageMap["preconditions"]))
	}
	preconditions, err := ParsePreconditions(preconditionsInterface)
	if err != nil {
		return nil, err
	}
	stage.Preconditions = preconditions
	delete(stageMap, "preconditions")

	if err := mapstructure.Decode(stageMap, stage); err != nil {
		return nil, err
	}
	return stage, nil
}
