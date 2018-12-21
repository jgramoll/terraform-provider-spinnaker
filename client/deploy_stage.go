package client

// DeployStageType deploy stage
var DeployStageType StageType = "deploy"

func init() {
	stageFactories[DeployStageType] = func() interface{} {
		return NewDeployStage()
	}
}

// StageEnabled when stage is enabled
type StageEnabled struct {
	Expression string `json:"expression"`
	Type       string `json:"type"`
}

// DeployStage for pipeline
type DeployStage struct {
	Name                 string    `json:"name"`
	RefID                string    `json:"refId"`
	Type                 StageType `json:"type"`
	RequisiteStageRefIds []string  `json:"requisiteStageRefIds"`

	CompleteOtherBranchesThenFail bool `json:"completeOtherBranchesThenFail"`
	ContinuePipeline              bool `json:"continuePipeline"`
	FailOnFailedExpressions       bool `json:"failOnFailedExpressions"`
	FailPipeline                  bool `json:"failPipeline"`

	Clusters                          []DeployStageCluster `json:"clusters"`
	OverrideTimeout                   bool                 `json:"overrideTimeout"`
	RestrictExecutionDuringTimeWindow bool                 `json:"restrictExecutionDuringTimeWindow"`
	RestrictedExecutionWindow         StageExecutionWindow `json:"restrictedExecutionWindow"`
	StageEnabled                      StageEnabled         `json:"stageEnabled"`
}

// NewDeployStage for pipeline
func NewDeployStage() *DeployStage {
	return &DeployStage{
		Type: DeployStageType,
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
