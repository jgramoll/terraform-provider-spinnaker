package client

import (
	"github.com/mitchellh/mapstructure"
)

// RunJobManifestStageType run job manifest stage
var RunJobManifestStageType StageType = "runJobManifest"

func init() {
	stageFactories[RunJobManifestStageType] = parseRunJobManifestStage
}

type RunJobManifestStage struct {
	BaseStage `mapstructure:",squash"`

	Account               string   `json:"account"`
	Alias                 string   `json:"alias"`
	Application           string   `json:"application"`
	CloudProvider         string   `json:"cloudProvider"`
	ConsumeArtifactSource string   `json:"consumeArtifactSource"`
	Credentials           string   `json:"credentails"`
	Manifest              Manifest `json:"manifest"`
}

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

	if err := mapstructure.Decode(stageMap, stage); err != nil {
		return nil, err
	}
	return stage, nil
}
