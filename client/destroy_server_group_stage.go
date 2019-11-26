package client

import (
	"github.com/mitchellh/mapstructure"
)

// DestroyServerGroupStageType destroy server group stage
var DestroyServerGroupStageType StageType = "destroyServerGroup"

func init() {
	stageFactories[DestroyServerGroupStageType] = parseDestroyServerGroupStage
}

// DestroyServerGroupStage for pipeline
type DestroyServerGroupStage struct {
	BaseStage `mapstructure:",squash"`

	CloudProvider     string   `json:"cloudProvider"`
	CloudProviderType string   `json:"cloudProviderType"`
	Cluster           string   `json:"cluster"`
	Credentials       string   `json:"credentials"`
	Moniker           *Moniker `json:"moniker"`
	Regions           []string `json:"regions"`
	Target            string   `json:"target"`
}

// NewDestroyServerGroupStage for pipeline
func NewDestroyServerGroupStage() *DestroyServerGroupStage {
	return &DestroyServerGroupStage{
		BaseStage: *newBaseStage(DestroyServerGroupStageType),
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
