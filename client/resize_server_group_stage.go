package client

import (
	"github.com/mitchellh/mapstructure"
)

// ResizeServerGroupStageType resize server group stage
var ResizeServerGroupStageType StageType = "resizeServerGroup"

func init() {
	stageFactories[ResizeServerGroupStageType] = parseResizeServerGroupStage
}

// ResizeServerGroupStage for pipeline
type serializableResizeServerGroupStage struct {
}

// ResizeServerGroupStage for pipeline
type ResizeServerGroupStage struct {
	BaseStage `mapstructure:",squash"`

	Action            string    `json:"action"`
	Capacity          *Capacity `json:"capacity"`
	CloudProvider     string    `json:"cloudProvider"`
	CloudProviderType string    `json:"cloudProviderType"`
	Cluster           string    `json:"cluster"`
	Credentials       string    `json:"credentials"`
	Moniker           *Moniker  `json:"moniker"`
	Regions           []string  `json:"regions"`
	ResizeType        string    `json:"resizeType"`
	Target            string    `json:"target"`

	TargetHealthyDeployPercentage int `json:"targetHealthyDeployPercentage"`
}

// NewResizeServerGroupStage for pipeline
func NewResizeServerGroupStage() *ResizeServerGroupStage {
	return &ResizeServerGroupStage{
		BaseStage: *newBaseStage(ResizeServerGroupStageType),
	}
}

func parseResizeServerGroupStage(stageMap map[string]interface{}) (Stage, error) {
	stage := NewResizeServerGroupStage()
	if err := stage.parseBaseStage(stageMap); err != nil {
		return nil, err
	}

	if err := mapstructure.Decode(stageMap, stage); err != nil {
		return nil, err
	}
	return stage, nil
}
