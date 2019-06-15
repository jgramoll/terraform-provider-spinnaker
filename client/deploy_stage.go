package client

import (
	"github.com/mitchellh/mapstructure"
)

// DeployStageType deploy stage
var DeployStageType StageType = "deploy"

func init() {
	stageFactories[DeployStageType] = parseDeployStage
}

// StageEnabled when stage is enabled
type StageEnabled struct {
	Expression string `json:"expression"`
	Type       string `json:"type"`
}

type serializableDeployStage struct {
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

	Clusters *[]*DeployStageCluster `json:"clusters"`
}

// DeployStage for pipeline
type DeployStage struct {
	*serializableDeployStage
	Notifications *[]*Notification `json:"notifications"`
}

func newSerializableDeployStage() *serializableDeployStage {
	return &serializableDeployStage{
		Type:                 DeployStageType,
		FailPipeline:         true,
		RequisiteStageRefIds: []string{},
	}
}

// NewDeployStage for pipeline
func NewDeployStage() *DeployStage {
	return &DeployStage{
		serializableDeployStage: newSerializableDeployStage(),
	}
}

// GetName for Stage interface
func (s *DeployStage) GetName() string {
	return s.Name
}

// GetType for Stage interface
func (s *DeployStage) GetType() StageType {
	return s.Type
}

// GetRefID for Stage interface
func (s *DeployStage) GetRefID() string {
	return s.RefID
}

func parseDeployStage(stageMap map[string]interface{}) (Stage, error) {
	stage := newSerializableDeployStage()
	if err := mapstructure.WeakDecode(stageMap, stage); err != nil {
		return nil, err
	}

	notifications, err := parseNotifications(stageMap["notifications"])
	if err != nil {
		return nil, err
	}
	return &DeployStage{
		serializableDeployStage: stage,
		Notifications:           notifications,
	}, nil
}
