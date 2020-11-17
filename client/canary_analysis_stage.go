package client

import (
	"github.com/mitchellh/mapstructure"
)

// CanaryAnalysisStageType stage type
var CanaryAnalysisStageType StageType = "kayentaCanary"

func init() {
	stageFactories[CanaryAnalysisStageType] = parseCanaryAnalysisStage
}

// CanaryAnalysisStage canary analysis stage
type CanaryAnalysisStage struct {
	BaseStage `mapstructure:",squash"`

	AnalysisType string                `json:"analysisType"`
	CanaryConfig *CanaryAnalysisConfig `json:"canaryConfig"`
	Deployments  *[]*DeploymentCluster `json:"deployments,omitempty"`
}

// NewCanaryAnalysisStage new canary analysis stage
func NewCanaryAnalysisStage() *CanaryAnalysisStage {
	return &CanaryAnalysisStage{
		BaseStage:    *newBaseStage(CanaryAnalysisStageType),
		CanaryConfig: NewCanaryAnalysisConfig(),
	}
}

func parseCanaryAnalysisStage(stageMap map[string]interface{}) (Stage, error) {
	stage := NewCanaryAnalysisStage()
	if err := stage.parseBaseStage(stageMap); err != nil {
		return nil, err
	}

	if err := mapstructure.Decode(stageMap, stage); err != nil {
		return nil, err
	}
	return stage, nil
}
