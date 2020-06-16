package client

import (
	"github.com/mitchellh/mapstructure"
)

// ScriptStageType Script stage
var ScriptStageType StageType = "script"

func init() {
	stageFactories[ScriptStageType] = parseScriptStage
}

// ScriptStage for pipeline
type ScriptStage struct {
	BaseStage `mapstructure:",squash"`

	Account           string   `json:"account"`
	Cluster           string   `json:"cluster"`
	Clusters          []string `json:"clusters"`
	Command           string   `json:"command"`
	Cmc               string   `json:"cmc"`
	Image             string   `json:"image"`
	PropertyFile      string   `json:"propertyFile"`
	Region            string   `json:"region"`
	Regions           []string `json:"regions"`
	RepoURL           string   `json:"repoUrl"`
	RepoBranch        string   `json:"repoBranch"`
	ScriptPath        string   `json:"scriptPath"`
	WaitForCompletion bool     `json:"waitForCompletion"`
}

// NewScriptStage for pipeline
func NewScriptStage() *ScriptStage {
	return &ScriptStage{
		BaseStage: *newBaseStage(ScriptStageType),
	}
}

func parseScriptStage(stageMap map[string]interface{}) (Stage, error) {
	stage := NewScriptStage()
	if err := stage.parseBaseStage(stageMap); err != nil {
		return nil, err
	}

	if err := mapstructure.Decode(stageMap, stage); err != nil {
		return nil, err
	}
	return stage, nil
}
