package client

import (
	"github.com/mitchellh/mapstructure"
)

// DisableManifestStageType delete manifest stage
var DisableManifestStageType StageType = "disableManifest"

func init() {
	stageFactories[DisableManifestStageType] = parseDisableManifestStage
}

// DisableManifestStage stage
type DisableManifestStage struct {
	BaseStage `mapstructure:",squash"`

	Account       string `json:"account"`
	App           string `json:"app"`
	CloudProvider string `json:"cloudProvider"`
	Cluster       string `json:"cluster"`
	Criteria      string `json:"criteria"`
	Kind          string `json:"kind"`
	// kinds          string `json:"kinds"`
	// labelSelectors string `json:"labelSelectors"`
	Location     string `json:"location"`
	ManifestName string `json:"manifestName,omitempty"`
	Mode         string `json:"mode"`
}

// NewDisableManifestStage new stage
func NewDisableManifestStage() *DisableManifestStage {
	return &DisableManifestStage{
		BaseStage: *newBaseStage(DisableManifestStageType),
	}
}

func parseDisableManifestStage(stageMap map[string]interface{}) (Stage, error) {
	stage := NewDisableManifestStage()
	if err := stage.parseBaseStage(stageMap); err != nil {
		return nil, err
	}

	if err := mapstructure.Decode(stageMap, stage); err != nil {
		return nil, err
	}
	return stage, nil
}
