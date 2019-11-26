package client

import (
	"github.com/mitchellh/mapstructure"
)

// DeployStageType deploy stage
var DeployStageType StageType = "deploy"

func init() {
	stageFactories[DeployStageType] = parseDeployStage
}

// StageEnabled when stage is enabled
type StageEnabled struct {
	Expression string `json:"expression"`
	Type       string `json:"type"`
}

// DeployStage for pipeline
type DeployStage struct {
	BaseStage `mapstructure:",squash"`

	Clusters *[]*DeploymentCluster `json:"clusters"`
}

// NewDeployStage for pipeline
func NewDeployStage() *DeployStage {
	return &DeployStage{
		BaseStage: *newBaseStage(DeployStageType),
	}
}

func parseDeployStage(stageMap map[string]interface{}) (Stage, error) {
	stage := NewDeployStage()
	if err := stage.parseBaseStage(stageMap); err != nil {
		return nil, err
	}

	if err := mapstructure.Decode(stageMap, stage); err != nil {
		return nil, err
	}
	return stage, nil
}
