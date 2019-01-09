package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

type pipelineStage struct {
	// TODO does this baseStage work?
	// baseStage
	Name                 string           `mapstructure:"name"`
	RefID                string           `mapstructure:"ref_id"`
	Type                 client.StageType `mapstructure:"type"`
	RequisiteStageRefIds []string         `mapstructure:"requisite_stage_ref_ids"`
	StageEnabled         []stageEnabled   `mapstructure:"stage_enabled"`

	Application                   string `mapstructure:"application"`
	CompleteOtherBranchesThenFail bool   `mapstructure:"complete_other_branches_then_fail"`

	ContinuePipeline   bool              `mapstructure:"continue_pipeline"`
	FailPipeline       bool              `mapstructure:"fail_pipeline"`
	Pipeline           string            `mapstructure:"target_pipeline"`
	PipelineParameters map[string]string `mapstructure:"pipeline_parameters"`
	WaitForCompletion  bool              `mapstructure:"wait_for_completion"`
}

func newPipelineStage() *pipelineStage {
	return &pipelineStage{Type: client.PipelineType}
}

func (s *pipelineStage) toClientStage() client.Stage {
	cs := client.NewPipelineStage()
	cs.Name = s.Name
	cs.RefID = s.RefID
	cs.RequisiteStageRefIds = s.RequisiteStageRefIds

	if len(s.StageEnabled) > 0 {
		newStageEnabled := client.StageEnabled(s.StageEnabled[0])
		cs.StageEnabled = &newStageEnabled
	}

	cs.Application = s.Application
	cs.CompleteOtherBranchesThenFail = s.CompleteOtherBranchesThenFail
	cs.ContinuePipeline = s.ContinuePipeline
	cs.FailPipeline = s.FailPipeline
	cs.Pipeline = s.Pipeline
	cs.PipelineParameters = s.PipelineParameters
	cs.WaitForCompletion = s.WaitForCompletion

	return cs
}

// TODO can we just update the ptr?
func (s *pipelineStage) fromClientStage(cs client.Stage) stage {
	clientStage := cs.(*client.PipelineStage)
	newStage := newPipelineStage()
	newStage.Name = clientStage.Name
	newStage.RefID = clientStage.RefID
	newStage.RequisiteStageRefIds = clientStage.RequisiteStageRefIds

	if clientStage.StageEnabled != nil {
		newStage.StageEnabled = append(newStage.StageEnabled, stageEnabled(*clientStage.StageEnabled))
	}

	newStage.Application = clientStage.Application
	newStage.CompleteOtherBranchesThenFail = clientStage.CompleteOtherBranchesThenFail
	newStage.ContinuePipeline = clientStage.ContinuePipeline
	newStage.FailPipeline = clientStage.FailPipeline
	newStage.Pipeline = clientStage.Pipeline
	newStage.PipelineParameters = clientStage.PipelineParameters
	newStage.WaitForCompletion = clientStage.WaitForCompletion

	return newStage
}

func (s *pipelineStage) SetResourceData(d *schema.ResourceData) {
	// TODO
	d.Set("name", s.Name)
	d.Set("requisite_stage_ref_ids", s.RequisiteStageRefIds)
}

func (s *pipelineStage) SetRefID(id string) {
	s.RefID = id
}

func (s *pipelineStage) GetRefID() string {
	return s.RefID
}
