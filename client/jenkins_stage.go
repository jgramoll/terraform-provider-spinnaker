package client

import (
	"github.com/mitchellh/mapstructure"
)

// JenkinsStageType jenkins stage
var JenkinsStageType StageType = "jenkins"

func init() {
	stageFactories[JenkinsStageType] = func(stageMap map[string]interface{}) (Stage, error) {
		stage := newSerializableJenkinsStage()
		if err := mapstructure.Decode(stageMap, stage); err != nil {
			return nil, err
		}

		notifications, err := parseNotifications(stageMap["notifications"])
		if err != nil {
			return nil, err
		}
		return &JenkinsStage{
			serializableJenkinsStage: stage,
			Notifications:            notifications,
		}, nil
	}
}

type serializableJenkinsStage struct {
	Name                 string    `json:"name"`
	RefID                string    `json:"refId"`
	Type                 StageType `json:"type"`
	RequisiteStageRefIds []string  `json:"requisiteStageRefIds"`
	SendNotifications    bool      `json:"sendNotifications"`

	CompleteOtherBranchesThenFail bool              `json:"completeOtherBranchesThenFail"`
	ContinuePipeline              bool              `json:"continuePipeline"`
	FailPipeline                  bool              `json:"failPipeline"`
	Job                           string            `json:"job"`
	MarkUnstableAsSuccessful      bool              `json:"markUnstableAsSuccessful"`
	Master                        string            `json:"master"`
	Parameters                    map[string]string `json:"parameters,omitempty"`
	PropertyFile                  string            `json:"propertyFile,omitempty"`
}

// JenkinsStage for pipeline
type JenkinsStage struct {
	*serializableJenkinsStage
	Notifications *[]*Notification `json:"notifications"`
}

// NewJenkinsStage for pipeline
func newSerializableJenkinsStage() *serializableJenkinsStage {
	return &serializableJenkinsStage{
		Type: JenkinsStageType,
	}
}

// NewJenkinsStage for pipeline
func NewJenkinsStage() *JenkinsStage {
	return &JenkinsStage{
		serializableJenkinsStage: newSerializableJenkinsStage(),
	}
}

// GetName for Stage interface
func (s *JenkinsStage) GetName() string {
	return s.Name
}

// GetType for Stage interface
func (s *JenkinsStage) GetType() StageType {
	return s.Type
}

// GetRefID for Stage interface
func (s *JenkinsStage) GetRefID() string {
	return s.RefID
}
