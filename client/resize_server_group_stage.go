package client

import (
	"github.com/mitchellh/mapstructure"
)

// ResizeServerGroupStageType resize server group stage
var ResizeServerGroupStageType StageType = "resizeServerGroup"

func init() {
	stageFactories[ResizeServerGroupStageType] = parseResizeServerGroupStage
}

// ResizeServerGroupStage for pipeline
type serializableResizeServerGroupStage struct {
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

	Action            string    `json:"action"`
	Capacity          *Capacity `json:"capacity"`
	CloudProvider     string    `json:"cloudProvider"`
	CloudProviderType string    `json:"cloudProviderType"`
	Cluster           string    `json:"cluster"`
	Credentials       string    `json:"credentials"`
	Moniker           *Moniker  `json:"moniker"`
	Regions           []string  `json:"regions"`
	ResizeType        string    `json:"resizeType"`
	Target            string    `json:"target"`

	TargetHealthyRollbackPercentage int `json:"targetHealthyRollbackPercentage"`
}

// ResizeServerGroupStage for pipeline
type ResizeServerGroupStage struct {
	*serializableResizeServerGroupStage
	Notifications *[]*Notification `json:"notifications"`
}

func newSerializableResizeServerGroupStage() *serializableResizeServerGroupStage {
	return &serializableResizeServerGroupStage{
		Type:                 ResizeServerGroupStageType,
		FailPipeline:         true,
		RequisiteStageRefIds: []string{},
	}
}

// NewResizeServerGroupStage for pipeline
func NewResizeServerGroupStage() *ResizeServerGroupStage {
	return &ResizeServerGroupStage{
		serializableResizeServerGroupStage: newSerializableResizeServerGroupStage(),
	}
}

// GetName for Stage interface
func (s *ResizeServerGroupStage) GetName() string {
	return s.Name
}

// GetType for Stage interface
func (s *ResizeServerGroupStage) GetType() StageType {
	return s.Type
}

// GetRefID for Stage interface
func (s *ResizeServerGroupStage) GetRefID() string {
	return s.RefID
}

func parseResizeServerGroupStage(stageMap map[string]interface{}) (Stage, error) {
	stage := newSerializableResizeServerGroupStage()
	if err := mapstructure.Decode(stageMap, stage); err != nil {
		return nil, err
	}

	notifications, err := parseNotifications(stageMap["notifications"])
	if err != nil {
		return nil, err
	}
	return &ResizeServerGroupStage{
		serializableResizeServerGroupStage: stage,
		Notifications:                      notifications,
	}, nil
}
