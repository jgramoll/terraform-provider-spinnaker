package client

import (
	"errors"
)

// ErrStageNotFound stage not found
var ErrStageNotFound = errors.New("Could not find stage")

// Stage interface for Pipeline stages
type Stage interface {
	GetName() string
	GetRefID() string
	GetType() StageType
}

// BaseStage attributes common to all Pipeline stages
type BaseStage struct {
	Name  string    `json:"name"`
	RefID string    `json:"refId"`
	Type  StageType `json:"type"`
}

func (pipeline *Pipeline) GetStage(stageID string) (Stage, error) {
	for _, s := range pipeline.Stages {
		if s.GetRefID() == stageID {
			return s, nil
		}
	}
	return nil, ErrStageNotFound
}

func (pipeline *Pipeline) UpdateStages(stage Stage) error {
	for i, pStage := range pipeline.Stages {
		if pStage.GetRefID() == stage.GetRefID() {
			pipeline.Stages[i] = stage
			return nil
		}
	}
	return ErrStageNotFound
}

func (pipeline *Pipeline) DeleteStage(stage Stage) error {
	for i, pStage := range pipeline.Stages {
		if pStage.GetRefID() == stage.GetRefID() {
			pipeline.Stages = append(pipeline.Stages[:i], pipeline.Stages[i+1:]...)
			return nil
		}
	}
	return ErrStageNotFound
}
