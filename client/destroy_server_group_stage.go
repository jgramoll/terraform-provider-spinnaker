package client

// DestroyServerGroupType destroy server group stage
var DestroyServerGroupType StageType = "destroyServerGroup"

func init() {
	stageFactories[DestroyServerGroupType] = func() interface{} {
		return NewDestroyServerGroupStage()
	}
}

// DestroyServerGroupStage for pipeline
type DestroyServerGroupStage struct {
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
	Target            string   `json:"target"`
}

// NewDestroyServerGroupStage for pipeline
func NewDestroyServerGroupStage() *DestroyServerGroupStage {
	return &DestroyServerGroupStage{
		// BaseStage: BaseStage{
		Type: DestroyServerGroupType,
		// },
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
