package client

import (
	"github.com/mitchellh/mapstructure"
)

// PatchManifestStageType bake manifest stage
var PatchManifestStageType StageType = "patchManifest"

func init() {
	stageFactories[PatchManifestStageType] = parsePatchManifestStage
}

// PatchManifestOptions options
type PatchManifestOptions struct {
	MergeStrategy string `json:"mergeStrategy"`
	Record        bool   `json:"record"`
}

// PatchManifestStage bake manifest
type PatchManifestStage struct {
	BaseStage `mapstructure:",squash"`

	Account       string `json:"account"`
	App           string `json:"app"`
	CloudProvider string `json:"cloudProvider"`
	Cluster       string `json:"cluster"`
	Criteria      string `json:"criteria"`
	Kind          string `json:"kind"`
	// kinds string `json:"kinds"`
	// labelSelectors string `json:"labelSelectors"`
	Location     string               `json:"location"`
	ManifestName string               `json:"manifestName"`
	Mode         string               `json:"mode"`
	Options      PatchManifestOptions `json:"options"`
	PatchBody    []string             `json:"patchBody"`
	Source       string               `json:"source"`
}

// NewPatchManifestStage bake manifest
func NewPatchManifestStage() *PatchManifestStage {
	return &PatchManifestStage{
		BaseStage: *newBaseStage(PatchManifestStageType),
	}
}

func parsePatchManifestStage(stageMap map[string]interface{}) (Stage, error) {
	stage := NewPatchManifestStage()
	if err := stage.parseBaseStage(stageMap); err != nil {
		return nil, err
	}

	if err := mapstructure.Decode(stageMap, stage); err != nil {
		return nil, err
	}
	return stage, nil
}
