package client

import (
	"log"

	"github.com/mitchellh/mapstructure"
)

type ParameterConfig struct{}

// PipelineWithoutStages deploy pipeline in application
type PipelineWithoutStages struct {
	Application          string            `json:"application"`
	Disabled             bool              `json:"disabled"`
	ID                   string            `json:"id"`
	Index                int               `json:"index"`
	KeepWaitingPipelines bool              `json:"keepWaitingPipelines"`
	LimitConcurrent      bool              `json:"limitConcurrent"`
	Name                 string            `json:"name"`
	Notifications        []Notification    `json:"notifications"`
	ParameterConfig      []ParameterConfig `json:"parameterConfig"`
	Triggers             []Trigger         `json:"triggers"`
	// TODO pointers?
}

// Pipeline deploy pipeline in application
type Pipeline struct {
	PipelineWithoutStages

	Stages []Stage `json:"stages"`
}

// NewPipelineWithoutStages Pipeline with default values
func NewPipelineWithoutStages() PipelineWithoutStages {
	return PipelineWithoutStages{
		Disabled:             false,
		KeepWaitingPipelines: false,
		LimitConcurrent:      true,
	}
}

// NewPipeline Pipeline with default values
func NewPipeline() *Pipeline {
	return &Pipeline{
		PipelineWithoutStages: NewPipelineWithoutStages(),
	}
}

func parsePipeline(pipelineHash map[string]interface{}) (*Pipeline, error) {
	pipelineWithoutStages := NewPipelineWithoutStages()
	if err := mapstructure.Decode(pipelineHash, &pipelineWithoutStages); err != nil {
		return nil, err
	}
	stagesHashInterface := pipelineHash["stages"]

	stages := []Stage{}
	if stagesHashInterface != nil {
		stagesToParse := stagesHashInterface.([]interface{})
		for _, stageInterface := range stagesToParse {
			stageMap := stageInterface.(map[string]interface{})

			stageTypeInterface, ok := stageMap["type"]
			if !ok {
				log.Println("[WARN] pipeline stage type is missing")
				continue
			}
			stageType := StageType(stageTypeInterface.(string))

			factory := stageFactories[stageType]
			if factory == nil {
				log.Printf("[WARN] unknown pipeline stage \"%s\"\n", stageType)
				continue
			}
			stage := factory()

			if err := mapstructure.Decode(stageMap, stage); err != nil {
				return nil, err
			}
			stages = append(stages, stage.(Stage))
		}
	}

	return &Pipeline{
		PipelineWithoutStages: pipelineWithoutStages,
		Stages:                stages,
	}, nil
}
