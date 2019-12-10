package client

// DestroyServerGroupStageType destroy server group stage
var DestroyServerGroupStageType StageType = "destroyServerGroup"

// DisableServerGroupStageType disable traffic to server group
var DisableServerGroupStageType StageType = "disableServerGroup"

func init() {
	stageFactories[DestroyServerGroupStageType] = parseDestroyServerGroupStage
	stageFactories[DisableServerGroupStageType] = parseDisableServerGroupStage
}

type serializableTargetServerGroupStage struct {
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

// TargetServerGroupStage for pipeline
type TargetServerGroupStage struct {
	*serializableTargetServerGroupStage
	Notifications *[]*Notification `json:"notifications"`
}

func newSerializableTargetServerGroupStage(stageType StageType) *serializableTargetServerGroupStage {
	return &serializableTargetServerGroupStage{
		Type:                 stageType,
		FailPipeline:         true,
		RequisiteStageRefIds: []string{},
	}
}

// NewTargetServerGroupStage for pipeline
func NewTargetServerGroupStage(stageType StageType) *TargetServerGroupStage {
	return &TargetServerGroupStage{
		serializableTargetServerGroupStage: newSerializableTargetServerGroupStage(stageType),
	}
}

// GetName for Stage interface
func (s *TargetServerGroupStage) GetName() string {
	return s.Name
}

// GetType for Stage interface
func (s *TargetServerGroupStage) GetType() StageType {
	return s.Type
}

// GetRefID for Stage interface
func (s *TargetServerGroupStage) GetRefID() string {
	return s.RefID
}
