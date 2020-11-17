package client

import (
	"github.com/mitchellh/mapstructure"
)

// ScaleManifestStageType scale manifest stage
var ScaleManifestStageType StageType = "scaleManifest"

func init() {
	stageFactories[ScaleManifestStageType] = parseScaleManifestStage
}

// ScaleManifestStage stage
type ScaleManifestStage struct {
	BaseStage `mapstructure:",squash"`

	Account        string                 `json:"account"`
	Application    string                 `json:"app"`
	CloudProvider  string                 `json:"cloudProvider"`
	Cluster        string                 `json:"cluster"`
	Criteria       string                 `json:"criteria"`
	Kind           string                 `json:"kind"`
	Kinds          []string               `json:"kinds,omitempty"`
	LabelSelectors map[string]interface{} `json:"labelSelectors,omitempty"`
	Location       string                 `json:"location"`
	ManifestName   string                 `json:"manifestName,omitempty"`
	Mode           string                 `json:"mode"`
	Replicas       string                 `json:"replicas"`
}

// NewScaleManifestStage new stage
func NewScaleManifestStage() *ScaleManifestStage {
	return &ScaleManifestStage{
		BaseStage: *newBaseStage(ScaleManifestStageType),
	}
}

func parseScaleManifestStage(stageMap map[string]interface{}) (Stage, error) {
	stage := NewScaleManifestStage()
	if err := stage.parseBaseStage(stageMap); err != nil {
		return nil, err
	}

	if err := mapstructure.Decode(stageMap, stage); err != nil {
		return nil, err
	}
	return stage, nil
}
