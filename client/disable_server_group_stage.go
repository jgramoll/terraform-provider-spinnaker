package client

import "github.com/mitchellh/mapstructure"

// DisableServerGroupStageType disable traffic to server group
var DisableServerGroupStageType StageType = "disableServerGroup"

func init() {
	stageFactories[DisableServerGroupStageType] = parseDisableServerGroupStage
}

// DisableServerGroupStage stage
type DisableServerGroupStage struct {
	BaseStage              `mapstructure:",squash"`
	TargetServerGroupStage `mapstructure:",squash"`
}

// NewDisableServerGroupStage new disable server group stage
func NewDisableServerGroupStage() *DisableServerGroupStage {
	return &DisableServerGroupStage{
		BaseStage:              *newBaseStage(DisableServerGroupStageType),
		TargetServerGroupStage: *newTargetServerGroupStage(),
	}
}

func parseDisableServerGroupStage(stageMap map[string]interface{}) (Stage, error) {
	stage := NewDisableServerGroupStage()
	if err := stage.parseBaseStage(stageMap); err != nil {
		return nil, err
	}

	if err := mapstructure.Decode(stageMap, stage); err != nil {
		return nil, err
	}
	return stage, nil
}
