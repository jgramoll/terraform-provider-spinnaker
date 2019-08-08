package client

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

// DeleteManifestStageType delete manifest stage
var DeleteManifestStageType StageType = "deleteManifest"

func init() {
	stageFactories[DeleteManifestStageType] = parseDeleteManifestStage
}

type DeleteManifestStage struct {
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

	Account       string                 `json:"account"`
	App           string                 `json:"app"`
	CloudProvider string                 `json:"cloudProvider"`
	Location      string                 `json:"location"`
	ManifestName  string                 `json:"manifestName"`
	Mode          DeleteManifestMode     `json:"mode"`
	Options       *DeleteManifestOptions `json:"options"`
}

func NewDeleteManifestStage() *DeleteManifestStage {
	return &DeleteManifestStage{
		Type:                 DeleteManifestStageType,
		FailPipeline:         true,
		RequisiteStageRefIds: []string{},
	}
}

// GetName for Stage interface
func (s *DeleteManifestStage) GetName() string {
	return s.Name
}

// GetType for Stage interface
func (s *DeleteManifestStage) GetType() StageType {
	return s.Type
}

// GetRefID for Stage interface
func (s *DeleteManifestStage) GetRefID() string {
	return s.RefID
}

func parseDeleteManifestStage(stageMap map[string]interface{}) (Stage, error) {
	notifications, err := parseNotifications(stageMap["notifications"])
	if err != nil {
		return nil, err
	}
	modeString, ok := stageMap["mode"].(string)
	if !ok {
		return nil, fmt.Errorf("Could not parse delete manifest mode %v", stageMap["mode"])
	}

	delete(stageMap, "notifications")
	delete(stageMap, "mode")
	stage := NewDeleteManifestStage()
	if err := mapstructure.Decode(stageMap, stage); err != nil {
		return nil, err
	}

	mode, err := ParseDeleteManifestMode(modeString)
	if err != nil {
		return nil, err
	}
	stage.Mode = mode
	stage.Notifications = notifications

	return stage, nil
}
