package client

import (
	"github.com/mitchellh/mapstructure"
)

// EnableManifestStageType delete manifest stage
var EnableManifestStageType StageType = "enableManifest"

func init() {
	stageFactories[EnableManifestStageType] = parseEnableManifestStage
}

// EnableManifestStage stage
type EnableManifestStage struct {
	BaseStage `mapstructure:",squash"`

	Account       string `json:"account"`
	App           string `json:"app"`
	CloudProvider string `json:"cloudProvider"`
	Cluster       string `json:"cluster"`
	Criteria      string `json:"criteria"`
	Kind          string `json:"kind"`
	Location      string `json:"location"`
	Mode          string `json:"mode"`
}

// NewEnableManifestStage new stage
func NewEnableManifestStage() *EnableManifestStage {
	return &EnableManifestStage{
		BaseStage: *newBaseStage(EnableManifestStageType),
	}
}

func parseEnableManifestStage(stageMap map[string]interface{}) (Stage, error) {
	stage := NewEnableManifestStage()
	if err := stage.parseBaseStage(stageMap); err != nil {
		return nil, err
	}

	if err := mapstructure.Decode(stageMap, stage); err != nil {
		return nil, err
	}
	return stage, nil
}
