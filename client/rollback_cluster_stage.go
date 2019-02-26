package client

import (
	"github.com/mitchellh/mapstructure"
)

// RollbackClusterStageType rollback cluster stage
var RollbackClusterStageType StageType = "rollbackCluster"

func init() {
	stageFactories[RollbackClusterStageType] = parseRollbackClusterStage
}

// RollbackClusterStage for pipeline
type serializableRollbackClusterStage struct {
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

	TargetHealthyRollbackPercentage int `json:"targetHealthyRollbackPercentage"`
}

// RollbackClusterStage for pipeline
type RollbackClusterStage struct {
	*serializableRollbackClusterStage
	Notifications *[]*Notification `json:"notifications"`
}

func newSerializableRollbackClusterStage() *serializableRollbackClusterStage {
	return &serializableRollbackClusterStage{
		Type:                 RollbackClusterStageType,
		FailPipeline:         true,
		RequisiteStageRefIds: []string{},
	}
}

// NewRollbackClusterStage for pipeline
func NewRollbackClusterStage() *RollbackClusterStage {
	return &RollbackClusterStage{
		serializableRollbackClusterStage: newSerializableRollbackClusterStage(),
	}
}

// GetName for Stage interface
func (s *RollbackClusterStage) GetName() string {
	return s.Name
}

// GetType for Stage interface
func (s *RollbackClusterStage) GetType() StageType {
	return s.Type
}

// GetRefID for Stage interface
func (s *RollbackClusterStage) GetRefID() string {
	return s.RefID
}

func parseRollbackClusterStage(stageMap map[string]interface{}) (Stage, error) {
	stage := newSerializableRollbackClusterStage()
	if err := mapstructure.Decode(stageMap, stage); err != nil {
		return nil, err
	}

	notifications, err := parseNotifications(stageMap["notifications"])
	if err != nil {
		return nil, err
	}
	return &RollbackClusterStage{
		serializableRollbackClusterStage: stage,
		Notifications:                    notifications,
	}, nil
}
