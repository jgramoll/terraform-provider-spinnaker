package client

import (
	"errors"
	"log"
)

// ErrStageNotFound stage not found
var ErrStageNotFound = errors.New("Could not find stage")

// Stage interface for Pipeline stages
type Stage interface {
	GetName() string
	GetRefID() string
	GetType() StageType
}

// GetStage get stage
func (pipeline *Pipeline) GetStage(stageID string) (Stage, error) {
	if pipeline.Stages != nil {
		for _, s := range *pipeline.Stages {
			if s.GetRefID() == stageID {
				return s, nil
			}
		}
	}
	return nil, ErrStageNotFound
}

// UpdateStage update stage
func (pipeline *Pipeline) UpdateStage(stage Stage) error {
	if pipeline.Stages != nil {
		for i, pStage := range *pipeline.Stages {
			if pStage.GetRefID() == stage.GetRefID() {
				(*pipeline.Stages)[i] = stage
				return nil
			}
		}
	}
	return ErrStageNotFound
}

// DeleteStage delete stage
func (pipeline *Pipeline) DeleteStage(stageID string) error {
	if pipeline.Stages != nil {
		stages := *pipeline.Stages
		for i, pStage := range stages {
			if pStage.GetRefID() == stageID {
				stages = append(stages[:i], stages[i+1:]...)
				pipeline.Stages = &stages
				return nil
			}
		}
	}
	return ErrStageNotFound
}

func parseStages(stagesHashInterface interface{}) (*[]Stage, error) {
	stages := []Stage{}
	if stagesHashInterface == nil {
		return &stages, nil
	}

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
		stage, err := factory(stageMap)
		if err != nil {
			return nil, err
		}
		stages = append(stages, stage)
	}
	return &stages, nil
}
