package client

import (
	"github.com/mitchellh/mapstructure"
)

// UndoRolloutManifestStageType undo rollout stage
var UndoRolloutManifestStageType StageType = "undoRolloutManifest"

func init() {
	stageFactories[UndoRolloutManifestStageType] = parseUndoRolloutManifestStage
}

// UndoRolloutManifestStage undo stage
type UndoRolloutManifestStage struct {
	BaseStage `mapstructure:",squash"`

	Account          string `json:"account"`
	CloudProvider    string `json:"cloudProvider"`
	Location         string `json:"location"`
	ManifestName     string `json:"manifestName,omitempty"`
	Mode             string `json:"mode"`
	NumRevisionsBack int    `json:"numRevisionsBack"`
}

// NewUndoRolloutManifestStage new stage
func NewUndoRolloutManifestStage() *UndoRolloutManifestStage {
	return &UndoRolloutManifestStage{
		BaseStage: *newBaseStage(UndoRolloutManifestStageType),
	}
}

func parseUndoRolloutManifestStage(stageMap map[string]interface{}) (Stage, error) {
	stage := NewUndoRolloutManifestStage()
	if err := stage.parseBaseStage(stageMap); err != nil {
		return nil, err
	}

	if err := mapstructure.Decode(stageMap, stage); err != nil {
		return nil, err
	}
	return stage, nil
}
