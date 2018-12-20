package client

// JenkinsStageType jenkins stage
var JenkinsStageType StageType = "jenkins"

func init() {
	stageFactories[JenkinsStageType] = func() interface{} {
		return NewJenkinsStage()
	}
}

// JenkinsStage for pipeline
type JenkinsStage struct {
	Name                 string    `json:"name"`
	RefID                string    `json:"refId"`
	Type                 StageType `json:"type"`
	RequisiteStageRefIds []string  `json:"requisiteStageRefIds"`

	CompleteOtherBranchesThenFail bool              `json:"completeOtherBranchesThenFail"`
	ContinuePipeline              bool              `json:"continuePipeline"`
	FailPipeline                  bool              `json:"failPipeline"`
	Job                           string            `json:"job"`
	MarkUnstableAsSuccessful      bool              `json:"markUnstableAsSuccessful"`
	Master                        string            `json:"master"`
	Parameters                    map[string]string `json:"parameters"`
	PropertyFile                  string            `json:"propertyFile"`
}

// NewJenkinsStage for pipeline
func NewJenkinsStage() *JenkinsStage {
	return &JenkinsStage{
		Type: JenkinsStageType,
	}
}

// GetName for Stage interface
func (s *JenkinsStage) GetName() string {
	return s.Name
}

// GetType for Stage interface
func (s *JenkinsStage) GetType() StageType {
	return s.Type
}

// GetRefID for Stage interface
func (s *JenkinsStage) GetRefID() string {
	return s.RefID
}
