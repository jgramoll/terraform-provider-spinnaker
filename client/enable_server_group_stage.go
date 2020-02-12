package client

import (
	"github.com/mitchellh/mapstructure"
)

// EnableServerGroupStageType enable server group stage
var EnableServerGroupStageType StageType = "enableServerGroup"

func init() {
	stageFactories[EnableServerGroupStageType] = parseEnableServerGroupStage
}

// EnableServerGroupStage enable server group stage
type EnableServerGroupStage struct {
	BaseStage `mapstructure:",squash"`

	CloudProvider                  string   `json:"cloudProvider"`
	CloudProviderType              string   `json:"cloudProviderType"`
	Cluster                        string   `json:"cluster"`
	Credentials                    string   `json:"credentials"`
	InterestingHealthProviderNames []string `json:"interestingHealthProviderNames"`
	Namespaces                     []string `json:"namespaces"`
	Regions                        []string `json:"regions"`
	Target                         string   `json:"target"`
}

// NewEnableServerGroupStage new enable server group stage
func NewEnableServerGroupStage() *EnableServerGroupStage {
	return &EnableServerGroupStage{
		BaseStage: *newBaseStage(EnableServerGroupStageType),
	}
}

func parseEnableServerGroupStage(stageMap map[string]interface{}) (Stage, error) {
	stage := NewEnableServerGroupStage()
	if err := stage.parseBaseStage(stageMap); err != nil {
		return nil, err
	}

	if err := mapstructure.Decode(stageMap, stage); err != nil {
		return nil, err
	}
	return stage, nil
}
