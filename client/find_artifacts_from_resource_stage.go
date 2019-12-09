package client

import (
	"github.com/mitchellh/mapstructure"
)

// FindArtifactsFromResourceStageType bake stage
var FindArtifactsFromResourceStageType StageType = "findArtifactsFromResource"

func init() {
	stageFactories[FindArtifactsFromResourceStageType] = parseFindArtifactsStage
}

// FindArtifactsFromResourceStage for pipeline
type FindArtifactsFromResourceStage struct {
	BaseStage `mapstructure:",squash"`

	Account       string `json:"account"`
	CloudProvider string `json:"cloudProvider"`
	Location      string `json:"location"`
	ManifestName  string `json:"manifestName"`
	Mode          string `json:"mode"`
}

// NewFindArtifactsFromResourceStage for pipeline
func NewFindArtifactsFromResourceStage() *FindArtifactsFromResourceStage {
	return &FindArtifactsFromResourceStage{
		BaseStage: *newBaseStage(FindArtifactsFromResourceStageType),
	}
}

func parseFindArtifactsStage(stageMap map[string]interface{}) (Stage, error) {
	stage := NewFindArtifactsFromResourceStage()
	if err := stage.parseBaseStage(stageMap); err != nil {
		return nil, err
	}

	if err := mapstructure.Decode(stageMap, stage); err != nil {
		return nil, err
	}
	return stage, nil
}
