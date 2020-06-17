package client

import (
	"github.com/mitchellh/mapstructure"
)

// BakeManifestStageType bake manifest stage
var BakeManifestStageType StageType = "bakeManifest"

func init() {
	stageFactories[BakeManifestStageType] = parseBakeManifestStage
}

// ManifestInputArtifact bake manifest stage
type ManifestInputArtifact struct {
	Account string `json:"account"`
	ID      string `json:"id"`
}

// BakeManifestStage bake manifest
type BakeManifestStage struct {
	BaseStage `mapstructure:",squash"`

	EvaluateOverrideExpressions bool `json:"evaluateOverrideExpressions"`

	InputArtifacts   []ManifestInputArtifact `json:"inputArtifacts"`
	Namespace        string                  `json:"namespace"`
	OutputName       string                  `json:"outputName"`
	Overrides        map[string]string       `json:"overrides"`
	RawOverrides     bool                    `json:"rawOverrides"`
	TemplateRenderer string                  `json:"templateRenderer"`
}

// NewBakeManifestStage bake manifest
func NewBakeManifestStage() *BakeManifestStage {
	return &BakeManifestStage{
		BaseStage: *newBaseStage(BakeManifestStageType),
	}
}

func parseBakeManifestStage(stageMap map[string]interface{}) (Stage, error) {
	stage := NewBakeManifestStage()
	if err := stage.parseBaseStage(stageMap); err != nil {
		return nil, err
	}

	if err := mapstructure.Decode(stageMap, stage); err != nil {
		return nil, err
	}
	return stage, nil
}
