package client

import (
	"github.com/mitchellh/mapstructure"
)

// RollbackClusterType rollback cluster stage
var RollbackClusterType StageType = "rollbackCluster"

func init() {
	stageFactories[RollbackClusterType] = func(stageMap map[string]interface{}) (Stage, error) {
		stage := NewRollbackClusterStage()
		if err := mapstructure.Decode(stageMap, stage); err != nil {
			return nil, err
		}
		return stage, nil
	}
}

// RollbackClusterStage for pipeline
type RollbackClusterStage struct {
	// TODO why does BaseStage not like mapstructure
	// BaseStage
	Name                 string        `json:"name"`
	RefID                string        `json:"refId"`
	Type                 StageType     `json:"type"`
	RequisiteStageRefIds []string      `json:"requisiteStageRefIds"`
	StageEnabled         *StageEnabled `json:"stageEnabled"`

	CloudProvider     string   `json:"cloudProvider"`
	CloudProviderType string   `json:"cloudProviderType"`
	Cluster           string   `json:"cluster"`
	Credentials       string   `json:"credentials"`
	Moniker           *Moniker `json:"moniker"`
	Regions           []string `json:"regions"`

	TargetHealthyRollbackPercentage int `json:"targetHealthyRollbackPercentage"`
}

// NewRollbackClusterStage for pipeline
func NewRollbackClusterStage() *RollbackClusterStage {
	return &RollbackClusterStage{
		// BaseStage: BaseStage{
		Type: RollbackClusterType,
		// },
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
