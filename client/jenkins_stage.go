package client

import (
	"github.com/mitchellh/mapstructure"
)

// JenkinsStageType jenkins stage
var JenkinsStageType StageType = "jenkins"

func init() {
	stageFactories[JenkinsStageType] = parseJenkinsStage
}

type serializableJenkinsStage struct {
}

// JenkinsStage for pipeline
type JenkinsStage struct {
	BaseStage `mapstructure:",squash"`

	Job                      string                 `json:"job"`
	MarkUnstableAsSuccessful bool                   `json:"markUnstableAsSuccessful"`
	Master                   string                 `json:"master"`
	Parameters               map[string]interface{} `json:"parameters,omitempty"`
	PropertyFile             string                 `json:"propertyFile,omitempty"`
}

// NewJenkinsStage for pipeline
func NewJenkinsStage() *JenkinsStage {
	return &JenkinsStage{
		BaseStage: *newBaseStage(JenkinsStageType),
	}
}

func parseJenkinsStage(stageMap map[string]interface{}) (Stage, error) {
	stage := NewJenkinsStage()
	if err := stage.parseBaseStage(stageMap); err != nil {
		return nil, err
	}

	if err := mapstructure.Decode(stageMap, stage); err != nil {
		return nil, err
	}
	return stage, nil
}
