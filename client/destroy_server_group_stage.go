package client

import (
	"github.com/mitchellh/mapstructure"
)

// DestroyServerGroupStageType destroy server group stage
var DestroyServerGroupStageType StageType = "destroyServerGroup"

func init() {
	stageFactories[DestroyServerGroupStageType] = parseDestroyServerGroupStage
}

type serializableDestroyServerGroupStage struct {
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

	CloudProvider     string   `json:"cloudProvider"`
	CloudProviderType string   `json:"cloudProviderType"`
	Cluster           string   `json:"cluster"`
	Credentials       string   `json:"credentials"`
	Moniker           *Moniker `json:"moniker"`
	Regions           []string `json:"regions"`
	Target            string   `json:"target"`
}

// DestroyServerGroupStage for pipeline
type DestroyServerGroupStage struct {
	*serializableDestroyServerGroupStage
	Notifications *[]*Notification `json:"notifications"`
}

func newSerializableDestroyServerGroupStage() *serializableDestroyServerGroupStage {
	return &serializableDestroyServerGroupStage{
		Type:                 DestroyServerGroupStageType,
		FailPipeline:         true,
		RequisiteStageRefIds: []string{},
	}
}

// NewDestroyServerGroupStage for pipeline
func NewDestroyServerGroupStage() *DestroyServerGroupStage {
	return &DestroyServerGroupStage{
		serializableDestroyServerGroupStage: newSerializableDestroyServerGroupStage(),
	}
}

// GetName for Stage interface
func (s *DestroyServerGroupStage) GetName() string {
	return s.Name
}

// GetType for Stage interface
func (s *DestroyServerGroupStage) GetType() StageType {
	return s.Type
}

// GetRefID for Stage interface
func (s *DestroyServerGroupStage) GetRefID() string {
	return s.RefID
}

func parseDestroyServerGroupStage(stageMap map[string]interface{}) (Stage, error) {
	stage := newSerializableDestroyServerGroupStage()
	if err := mapstructure.Decode(stageMap, stage); err != nil {
		return nil, err
	}

	notifications, err := parseNotifications(stageMap["notifications"])
	if err != nil {
		return nil, err
	}
	return &DestroyServerGroupStage{
		serializableDestroyServerGroupStage: stage,
		Notifications:                       notifications,
	}, nil
}
