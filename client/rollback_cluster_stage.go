package client

import (
	"github.com/mitchellh/mapstructure"
)

// RollbackClusterStageType rollback cluster stage
var RollbackClusterStageType StageType = "rollbackCluster"

func init() {
	stageFactories[RollbackClusterStageType] = parseRollbackClusterStage
}

// RollbackClusterStage for pipeline
type RollbackClusterStage struct {
	BaseStage `mapstructure:",squash"`

	CloudProvider     string   `json:"cloudProvider"`
	CloudProviderType string   `json:"cloudProviderType"`
	Cluster           string   `json:"cluster"`
	Credentials       string   `json:"credentials"`
	Moniker           *Moniker `json:"moniker"`
	Regions           []string `json:"regions"`

	TargetHealthyRollbackPercentage int `json:"targetHealthyRollbackPercentage"`
}

// NewRollbackClusterStage for pipeline
func NewRollbackClusterStage() *RollbackClusterStage {
	return &RollbackClusterStage{
		BaseStage: *newBaseStage(RollbackClusterStageType),
	}
}

func parseRollbackClusterStage(stageMap map[string]interface{}) (Stage, error) {
	stage := NewRollbackClusterStage()
	if err := stage.parseBaseStage(stageMap); err != nil {
		return nil, err
	}

	if err := mapstructure.Decode(stageMap, stage); err != nil {
		return nil, err
	}
	return stage, nil
}
