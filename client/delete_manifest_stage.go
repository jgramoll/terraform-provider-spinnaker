package client

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

// DeleteManifestStageType delete manifest stage
var DeleteManifestStageType StageType = "deleteManifest"

func init() {
	stageFactories[DeleteManifestStageType] = parseDeleteManifestStage
}

// DeleteManifestStage stage
type DeleteManifestStage struct {
	BaseStage `mapstructure:",squash"`

	Account       string                 `json:"account"`
	App           string                 `json:"app"`
	CloudProvider string                 `json:"cloudProvider"`
	Location      string                 `json:"location"`
	ManifestName  string                 `json:"manifestName"`
	Mode          DeleteManifestMode     `json:"mode"`
	Options       *DeleteManifestOptions `json:"options"`
}

// NewDeleteManifestStage new stage
func NewDeleteManifestStage() *DeleteManifestStage {
	return &DeleteManifestStage{
		BaseStage: *newBaseStage(DeleteManifestStageType),
	}
}

func parseDeleteManifestStage(stageMap map[string]interface{}) (Stage, error) {
	stage := NewDeleteManifestStage()
	if err := stage.parseBaseStage(stageMap); err != nil {
		return nil, err
	}

	modeString, ok := stageMap["mode"].(string)
	if !ok {
		return nil, fmt.Errorf("Could not parse delete manifest mode %v", stageMap["mode"])
	}
	mode, err := ParseDeleteManifestMode(modeString)
	if err != nil {
		return nil, err
	}
	stage.Mode = mode
	delete(stageMap, "mode")

	if err := mapstructure.Decode(stageMap, stage); err != nil {
		return nil, err
	}
	return stage, nil
}
