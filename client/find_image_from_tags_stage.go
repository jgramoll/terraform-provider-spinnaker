package client

import (
	"github.com/mitchellh/mapstructure"
)

// FindImageFromTagsStageType bake stage
var FindImageFromTagsStageType StageType = "findImageFromTags"

func init() {
	stageFactories[FindImageFromTagsStageType] = parseFindImageStage
}

type serializableFindImageFromTagsStage struct {
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

// FindImageFromTagsStage for pipeline
type FindImageFromTagsStage struct {
	*serializableFindImageFromTagsStage
	Notifications *[]*Notification `json:"notifications"`
}

func newserializableFindImageStage() *serializableFindImageFromTagsStage {
	return &serializableFindImageFromTagsStage{
		Type:                 FindImageFromTagsStageType,
		FailPipeline:         true,
		RequisiteStageRefIds: []string{},
	}
}

// NewFindImageStage for pipeline
func NewFindImageStage() *FindImageFromTagsStage {
	return &FindImageFromTagsStage{
		serializableFindImageFromTagsStage: newserializableFindImageStage(),
	}
}

// GetName for Stage interface
func (s *FindImageFromTagsStage) GetName() string {
	return s.Name
}

// GetType for Stage interface
func (s *FindImageFromTagsStage) GetType() StageType {
	return s.Type
}

// GetRefID for Stage interface
func (s *FindImageFromTagsStage) GetRefID() string {
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
	return &FindImageFromTagsStage{
		serializableFindImageFromTagsStage: stage,
		Notifications:                      notifications,
	}, nil
}
