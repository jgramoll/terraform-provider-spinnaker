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
}

// PipelineStage for pipeline
type PipelineStage struct {
	BaseStage `mapstructure:",squash"`

	Application        string                 `json:"application"`
	Pipeline           string                 `json:"pipeline"`
	PipelineParameters map[string]interface{} `json:"pipelineParameters"`
	WaitForCompletion  bool                   `json:"waitForCompletion"`
}

// NewPipelineStage for pipeline
func NewPipelineStage() *PipelineStage {
	return &PipelineStage{
		BaseStage: *newBaseStage(PipelineStageType),
	}
}

func parsePipelineStage(stageMap map[string]interface{}) (Stage, error) {
	stage := NewPipelineStage()
	if err := stage.parseBaseStage(stageMap); err != nil {
		return nil, err
	}

	if err := mapstructure.Decode(stageMap, stage); err != nil {
		return nil, err
	}
	return stage, nil
}
