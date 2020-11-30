package client

import "github.com/mitchellh/mapstructure"

// DestroyServerGroupStageType destroy server group stage
var DestroyServerGroupStageType StageType = "destroyServerGroup"

func init() {
	stageFactories[DestroyServerGroupStageType] = parseDestroyServerGroupStage
}

// DestroyServerGroupStage stage
type DestroyServerGroupStage struct {
	BaseStage              `mapstructure:",squash"`
	TargetServerGroupStage `mapstructure:",squash"`
}

// NewDestroyServerGroupStage new disable server group stage
func NewDestroyServerGroupStage() *DestroyServerGroupStage {
	return &DestroyServerGroupStage{
		BaseStage:              *newBaseStage(DestroyServerGroupStageType),
		TargetServerGroupStage: *newTargetServerGroupStage(),
	}
}

func parseDestroyServerGroupStage(stageMap map[string]interface{}) (Stage, error) {
	stage := NewDestroyServerGroupStage()
	if err := stage.parseBaseStage(stageMap); err != nil {
		return nil, err
	}

	if err := mapstructure.Decode(stageMap, stage); err != nil {
		return nil, err
	}
	return stage, nil
}
