package client

import (
	"github.com/mitchellh/mapstructure"
)

// PipelineStageType pipeline stage
var PipelineStageType StageType = "pipeline"

func init() {
	stageFactories[PipelineStageType] = parsePipelineStage
}

type serializablePipelineStage struct {
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

	Application        string            `json:"application"`
	Pipeline           string            `json:"pipeline"`
	PipelineParameters map[string]string `json:"pipelineParameters"`
	WaitForCompletion  bool              `json:"waitForCompletion"`
}

// PipelineStage for pipeline
type PipelineStage struct {
	*serializablePipelineStage
	Notifications *[]*Notification `json:"notifications"`
}

func newSerializablePipelineStage() *serializablePipelineStage {
	return &serializablePipelineStage{
		Type: PipelineStageType,
	}
}

// NewPipelineStage for pipeline
func NewPipelineStage() *PipelineStage {
	return &PipelineStage{
		serializablePipelineStage: newSerializablePipelineStage(),
	}
}

// GetName for Stage interface
func (s *PipelineStage) GetName() string {
	return s.Name
}

// GetType for Stage interface
func (s *PipelineStage) GetType() StageType {
	return s.Type
}

// GetRefID for Stage interface
func (s *PipelineStage) GetRefID() string {
	return s.RefID
}

func parsePipelineStage(stageMap map[string]interface{}) (Stage, error) {
	stage := newSerializablePipelineStage()
	if err := mapstructure.Decode(stageMap, stage); err != nil {
		return nil, err
	}

	notifications, err := parseNotifications(stageMap["notifications"])
	if err != nil {
		return nil, err
	}
	return &PipelineStage{
		serializablePipelineStage: stage,
		Notifications:             notifications,
	}, nil
}
