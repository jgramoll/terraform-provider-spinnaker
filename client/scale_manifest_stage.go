package client

import (
 	"github.com/mitchellh/mapstructure"
)

// ScaleManifestStageType scale manifest stage
var ScaleManifestStageType StageType = "scaleManifest"

func init() {
	stageFactories[ScaleManifestStageType] = parseScaleManifestStage
}

type ScaleManifestStage struct {
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

	Account        string                 `json:"account"`
	Application    string                 `json:"app"`
	CloudProvider  string                 `json:"cloudProvider"`
	Cluster        string                 `json:"cluster"`
	Criteria       string                 `json:"criteria"`
	IsNew          bool                   `json:"isNew"`
	Kind           string                 `json:"kind"`
	Kinds          []string               `json:"kinds,omitempty"`
	LabelSelectors map[string]interface{} `json:"labelSelectors,omitempty"`
	Location       string                 `json:"location"`
	ManifestName   string                 `json:"manifestName,omitempty"`
	Mode           string                 `json:"mode"`
	Replicas       string                 `json:"replicas"`
}

func NewScaleManifestStage() *ScaleManifestStage {
	return &ScaleManifestStage{
		Type:                 ScaleManifestStageType,
		FailPipeline:         true,
		RequisiteStageRefIds: []string{},
        // TODO defaults
	}
}

// GetName for Stage interface
func (s *ScaleManifestStage) GetName() string {
	return s.Name
}

// GetType for Stage interface
func (s *ScaleManifestStage) GetType() StageType {
	return s.Type
}

// GetRefID for Stage interface
func (s *ScaleManifestStage) GetRefID() string {
	return s.RefID
}

func parseScaleManifestStage(stageMap map[string]interface{}) (Stage, error) {
	stage := NewScaleManifestStage()
	notifications, err := parseNotifications(stageMap["notifications"])
	if err != nil {
		return nil, err
	}

	// Have to parse these seperate
	delete(stageMap, "notifications")
	if err := mapstructure.Decode(stageMap, stage); err != nil {
		return nil, err
	}

	stage.Notifications = notifications

	return stage, nil
}
