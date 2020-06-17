package client

// BaseStage all stages should have these properties
type BaseStage struct {
	Name                 string        `json:"name"`
	RefID                string        `json:"refId"`
	Type                 StageType     `json:"type"`
	RequisiteStageRefIds []string      `json:"requisiteStageRefIds"`
	SendNotifications    bool          `json:"sendNotifications"`
	StageEnabled         *StageEnabled `json:"stageEnabled"`
	Comments             string        `json:"comments,omitempty"`

	CompleteOtherBranchesThenFail     bool `json:"completeOtherBranchesThenFail"`
	ContinuePipeline                  bool `json:"continuePipeline"`
	FailOnFailedExpressions           bool `json:"failOnFailedExpressions"`
	FailPipeline                      bool `json:"failPipeline"`
	OverrideTimeout                   bool `json:"overrideTimeout"`
	StageTimeoutMS                    int  `json:"stageTimeoutMs,omitempty"`
	RestrictExecutionDuringTimeWindow bool `json:"restrictExecutionDuringTimeWindow"`

	RestrictedExecutionWindow *StageExecutionWindow `json:"restrictedExecutionWindow"`
	Notifications             *[]*Notification      `json:"notifications"`
	ExpectedArtifacts         *[]*ExpectedArtifact  `json:"expectedArtifacts"`
}

func newBaseStage(t StageType) *BaseStage {
	return &BaseStage{
		Type:                 t,
		FailPipeline:         true,
		RequisiteStageRefIds: []string{},
		ExpectedArtifacts:    &[]*ExpectedArtifact{},
	}
}

// GetName for Stage interface
func (s *BaseStage) GetName() string {
	return s.Name
}

// GetType for Stage interface
func (s *BaseStage) GetType() StageType {
	return s.Type
}

// GetRefID for Stage interface
func (s *BaseStage) GetRefID() string {
	return s.RefID
}

func (s *BaseStage) parseBaseStage(stageMap map[string]interface{}) error {
	notifications, err := parseNotifications(stageMap["notifications"])
	if err != nil {
		return err
	}
	s.Notifications = notifications
	delete(stageMap, "notifications")

	return nil
}
