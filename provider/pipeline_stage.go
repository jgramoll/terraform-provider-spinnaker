package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

type stage interface {
	fromClientStage(client.Stage) stage
	toClientStage() client.Stage
	SetResourceData(*schema.ResourceData)
	SetRefID(string)
	GetRefID() string
}

type baseStage struct {
	Name  string           `mapstructure:"name"`
	RefID string           `mapstructure:"ref_id"`
	Type  client.StageType `mapstructure:"type"`
}

// TODO make (s *Pipeline)
func getStage(stages []client.Stage, stageID string) (client.Stage, error) {
	for _, s := range stages {
		if s.GetRefID() == stageID {
			return s, nil
		}
	}
	return nil, ErrStageNotFound
}

// TODO make (s *Pipeline)
func updateStages(pipeline *client.Pipeline, stage client.Stage) error {
	for i, pStage := range pipeline.Stages {
		if pStage.GetRefID() == stage.GetRefID() {
			pipeline.Stages[i] = stage
			return nil
		}
	}
	return ErrStageNotFound
}

// TODO make (s *Pipeline)
func deleteStage(pipeline *client.Pipeline, stage client.Stage) error {
	for i, pStage := range pipeline.Stages {
		if pStage.GetRefID() == stage.GetRefID() {
			pipeline.Stages = append(pipeline.Stages[:i], pipeline.Stages[i+1:]...)
			return nil
		}
	}
	return ErrStageNotFound
}
