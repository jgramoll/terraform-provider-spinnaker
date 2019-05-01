package client

import (
	"github.com/mitchellh/mapstructure"
)

// FindImageStageType bake stage
var FindImageStageType StageType = "findImageFromTags"

func init() {
	stageFactories[FindImageStageType] = parseFindImageStage
}

type serializableFindImageStage struct {
	// BaseStage
	Name                              string                `json:"name"`
	RefID                             string                `json:"refId"`
	Type                              StageType             `json:"type"`
	RequisiteStageRefIds              []string              `json:"requisiteStageRefIds"`
	SendNotifications                 bool                  `json:"sendNotifications"`
	StageEnabled                      *StageEnabled         `json:"stageEnabled"`
	CompleteOtherBranchesThenFail     bool                  `json:"completeOtherBranchesThenFail"`
	ContinuePipeline                  bool                  `json:"continuePipeline"`
	FailOnFailedExpressions           bool                  `json:"failOnFailedExpressions"`
	FailPipeline                      bool                  `json:"failPipeline"`
	OverrideTimeout                   bool                  `json:"overrideTimeout"`
	RestrictExecutionDuringTimeWindow bool                  `json:"restrictExecutionDuringTimeWindow"`
	RestrictedExecutionWindow         *StageExecutionWindow `json:"restrictedExecutionWindow"`
	// End BaseStage

	CloudProvider     string            `json:"cloudProvider"`
	CloudProviderType string            `json:"cloudProviderType"`
	PackageName       string            `json:"packageName"`
	Regions           []string          `json:"regions"`
	Tags              map[string]string `json:"tags"`
}

// FindImageStage for pipeline
type FindImageStage struct {
	*serializableFindImageStage
	Notifications *[]*Notification `json:"notifications"`
}

func newserializableFindImageStage() *serializableFindImageStage {
	return &serializableFindImageStage{
		Type:                 FindImageStageType,
		FailPipeline:         true,
		RequisiteStageRefIds: []string{},
	}
}

// NewFindImageStage for pipeline
func NewFindImageStage() *FindImageStage {
	return &FindImageStage{
		serializableFindImageStage: newserializableFindImageStage(),
	}
}

// GetName for Stage interface
func (s *FindImageStage) GetName() string {
	return s.Name
}

// GetType for Stage interface
func (s *FindImageStage) GetType() StageType {
	return s.Type
}

// GetRefID for Stage interface
func (s *FindImageStage) GetRefID() string {
	return s.RefID
}

func parseFindImageStage(stageMap map[string]interface{}) (Stage, error) {
	stage := newserializableFindImageStage()
	if err := mapstructure.Decode(stageMap, stage); err != nil {
		return nil, err
	}

	notifications, err := parseNotifications(stageMap["notifications"])
	if err != nil {
		return nil, err
	}
	return &FindImageStage{
		serializableFindImageStage: stage,
		Notifications:              notifications,
	}, nil
}
