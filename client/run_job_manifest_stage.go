package client

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

// RunJobManifestStageType run job manifest stage
var RunJobManifestStageType StageType = "runJobManifest"

func init() {
	stageFactories[RunJobManifestStageType] = parseRunJobManifestStage
}

// RunJobManifestStage run job manifest stage
type RunJobManifestStage struct {
	BaseStage `mapstructure:",squash"`

	Account               string   `json:"account"`
	Alias                 string   `json:"alias"`
	Application           string   `json:"application"`
	CloudProvider         string   `json:"cloudProvider"`
	ConsumeArtifactSource string   `json:"consumeArtifactSource"`
	Credentials           string   `json:"credentials"`
	Manifest              Manifest `json:"manifest"`
	PropertyFile          string   `json:"propertyFile"`
	Source                string   `json:"source"`
}

// NewRunJobManifestStage new RunJobManifestStage
func NewRunJobManifestStage() *RunJobManifestStage {
	return &RunJobManifestStage{
		BaseStage: *newBaseStage(RunJobManifestStageType),

		Alias: "runJob",
	}
}

func parseRunJobManifestStage(stageMap map[string]interface{}) (Stage, error) {
	stage := NewRunJobManifestStage()
	if err := stage.parseBaseStage(stageMap); err != nil {
		return nil, err
	}

	manifestInterface, ok := stageMap["manifest"].(interface{})
	if !ok {
		return nil, fmt.Errorf("Could not parse run job manifest: %v", stageMap["manifest"])
	}
	manifest, err := ParseManifest(manifestInterface)
	if err != nil {
		return nil, err
	}
	stage.Manifest = manifest
	delete(stageMap, "manifest")

	if err := mapstructure.Decode(stageMap, stage); err != nil {
		return nil, err
	}
	return stage, nil
}
