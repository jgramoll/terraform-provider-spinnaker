package client

import (
	"github.com/mitchellh/mapstructure"
)

// ManualJudgmentStageType deploy stage
var ManualJudgmentStageType StageType = "manualJudgment"

func init() {
	stageFactories[ManualJudgmentStageType] = parseManualJudgmentStage
}

// JudgmentInputs inputs for judgment
type JudgmentInputs struct {
	Value string `json:"value"`
}

type serializableManualJudgmentStage struct {
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

	Instructions   string           `json:"instructions"`
	JudgmentInputs []JudgmentInputs `json:"judgmentInputs"`
}

// ManualJudgmentStage for pipeline
type ManualJudgmentStage struct {
	*serializableManualJudgmentStage
	Notifications *[]*Notification `json:"notifications"`
}

func newSerializableManualJudgmentStage() *serializableManualJudgmentStage {
	return &serializableManualJudgmentStage{
		Type:                 ManualJudgmentStageType,
		FailPipeline:         true,
		RequisiteStageRefIds: []string{},
	}
}

// NewManualJudgmentStage for pipeline
func NewManualJudgmentStage() *ManualJudgmentStage {
	return &ManualJudgmentStage{
		serializableManualJudgmentStage: newSerializableManualJudgmentStage(),
	}
}

// GetName for Stage interface
func (s *ManualJudgmentStage) GetName() string {
	return s.Name
}

// GetType for Stage interface
func (s *ManualJudgmentStage) GetType() StageType {
	return s.Type
}

// GetRefID for Stage interface
func (s *ManualJudgmentStage) GetRefID() string {
	return s.RefID
}

func parseManualJudgmentStage(stageMap map[string]interface{}) (Stage, error) {
	stage := newSerializableManualJudgmentStage()
	if err := mapstructure.Decode(stageMap, stage); err != nil {
		return nil, err
	}

	notifications, err := parseNotifications(stageMap["notifications"])
	if err != nil {
		return nil, err
	}
	return &ManualJudgmentStage{
		serializableManualJudgmentStage: stage,
		Notifications:                   notifications,
	}, nil
}
