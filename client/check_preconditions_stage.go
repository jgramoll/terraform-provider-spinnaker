package client

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/mitchellh/mapstructure"
)

// CheckPreconditionsStageType check preconditions stage
var CheckPreconditionsStageType StageType = "checkPreconditions"

func init() {
	stageFactories[CheckPreconditionsStageType] = parseCheckPreconditionsStage
}

type CheckPreconditionsStage struct {
	BaseStage `mapstructure:",squash"`

	Preconditions []Precondition `json:"preconditions"`
}

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

func ParsePreconditions(in []interface{}) ([]Precondition, error) {
	preconditions := []Precondition{}

	for _, preconditionInterface := range in {
		preconditionMap, ok := preconditionInterface.(map[string]interface{})
		if !ok {
			return nil, errors.New("invalid precondition")
		}
		typeString, ok := preconditionMap["type"].(string)
		if !ok {
			return nil, errors.New("missing or invalid precondition type")
		}
		preconditionType := PreconditionType(typeString)
		preconditionFunc, ok := preconditionFactories[preconditionType]
		if !ok {
			return nil, fmt.Errorf("unknown precondition %s", typeString)
		}
		precondition, err := preconditionFunc(preconditionMap)
		if err != nil {
			return nil, err
		}
		preconditions = append(preconditions, precondition)
	}

	return preconditions, nil
}

// func ParseDeployManifests(manifestInterface []interface{}) (*DeployManifests, error) {
// 	manifests := NewDeployManifests()
// 	for _, manifest := range manifestInterface {
// 		b, err := yaml.Marshal(manifest)
// 		if err != nil {
// 			return nil, err
// 		}
// 		*manifests = append(*manifests, string(b))
// 	}
// 	return manifests, nil
// }
