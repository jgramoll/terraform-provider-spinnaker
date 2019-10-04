package client

import (
	"github.com/mitchellh/mapstructure"
)

var CanaryAnalysisStageType StageType = "kayentaCanary"

func init() {
	stageFactories[CanaryAnalysisStageType] = parseCanaryAnalysisStage
}

type CanaryAnalysisStage struct {
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
	Notifications                     *[]*Notification      `json:"notifications"`
	// End BaseStage

	AnalysisType string                 `json:"analysisType"`
	CanaryConfig *CanaryAnalysisConfig  `json:"canaryConfig"`
	Deployments  *[]*DeployStageCluster `json:"deployments,omitempty"`
}

func NewCanaryAnalysisStage() *CanaryAnalysisStage {
	return &CanaryAnalysisStage{
		Type:                 CanaryAnalysisStageType,
		FailPipeline:         true,
		RequisiteStageRefIds: []string{},
		CanaryConfig:         NewCanaryAnalysisConfig(),
	}
}

// GetName for Stage interface
func (s *CanaryAnalysisStage) GetName() string {
	return s.Name
}

// GetType for Stage interface
func (s *CanaryAnalysisStage) GetType() StageType {
	return s.Type
}

// GetRefID for Stage interface
func (s *CanaryAnalysisStage) GetRefID() string {
	return s.RefID
}

func parseCanaryAnalysisStage(stageMap map[string]interface{}) (Stage, error) {
	stage := NewCanaryAnalysisStage()
	notifications, err := parseNotifications(stageMap["notifications"])
	if err != nil {
		return nil, err
	}
	delete(stageMap, "notifications")

	if err := mapstructure.Decode(stageMap, stage); err != nil {
		return nil, err
	}
	stage.Notifications = notifications
	return stage, nil
}
