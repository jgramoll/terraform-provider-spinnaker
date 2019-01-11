package client

import (
	"github.com/mitchellh/mapstructure"
)

// PipelineType pipeline stage
var PipelineType StageType = "pipeline"

func init() {
	stageFactories[PipelineType] = func(stageMap map[string]interface{}) (Stage, error) {
		stage := NewPipelineStage()
		if err := mapstructure.Decode(stageMap, stage); err != nil {
			return nil, err
		}
		return stage, nil
	}
}

// PipelineStage for pipeline
type PipelineStage struct {
	// TODO why does BaseStage not like mapstructure
	// BaseStage
	Name                 string        `json:"name"`
	RefID                string        `json:"refId"`
	Type                 StageType     `json:"type"`
	RequisiteStageRefIds []string      `json:"requisiteStageRefIds"`
	StageEnabled         *StageEnabled `json:"stageEnabled"`

	Application                   string            `json:"application"`
	CompleteOtherBranchesThenFail bool              `json:"completeOtherBranchesThenFail"`
	ContinuePipeline              bool              `json:"continuePipeline"`
	FailPipeline                  bool              `json:"failPipeline"`
	Pipeline                      string            `json:"pipeline"`
	PipelineParameters            map[string]string `json:"pipelineParameters"`
	WaitForCompletion             bool              `json:"waitForCompletion"`
}

// NewPipelineStage for pipeline
func NewPipelineStage() *PipelineStage {
	return &PipelineStage{
		// BaseStage: BaseStage{
		Type: PipelineType,
		// },
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
