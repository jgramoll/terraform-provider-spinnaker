package provider

import (
	"github.com/get-bridge/terraform-provider-spinnaker/client"
	"github.com/hashicorp/terraform/helper/schema"
)

type pipelineStage struct {
	baseStage `mapstructure:",squash"`

	Application        string                 `mapstructure:"application"`
	Pipeline           string                 `mapstructure:"target_pipeline"`
	PipelineParameters map[string]interface{} `mapstructure:"pipeline_parameters"`
	WaitForCompletion  bool                   `mapstructure:"wait_for_completion"`
}

func newPipelineStage() *pipelineStage {
	return &pipelineStage{
		baseStage: *newBaseStage(),
	}
}

func (s *pipelineStage) toClientStage(config *client.Config, refID string) (client.Stage, error) {
	cs := client.NewPipelineStage()
	err := s.baseToClientStage(&cs.BaseStage, refID, newDefaultNotificationInterface)
	if err != nil {
		return nil, err
	}

	cs.Application = s.Application
	cs.Pipeline = s.Pipeline
	cs.PipelineParameters = s.PipelineParameters
	cs.WaitForCompletion = s.WaitForCompletion

	return cs, nil
}

func (*pipelineStage) fromClientStage(cs client.Stage) (stage, error) {
	clientStage := cs.(*client.PipelineStage)
	newStage := newPipelineStage()
	err := newStage.baseFromClientStage(&clientStage.BaseStage, newDefaultNotificationInterface)
	if err != nil {
		return nil, err
	}

	newStage.Application = clientStage.Application
	newStage.Pipeline = clientStage.Pipeline
	newStage.PipelineParameters = clientStage.PipelineParameters
	newStage.WaitForCompletion = clientStage.WaitForCompletion

	return newStage, nil
}

func (s *pipelineStage) SetResourceData(d *schema.ResourceData) error {
	err := s.baseSetResourceData(d)
	if err != nil {
		return err
	}

	err = d.Set("application", s.Application)
	if err != nil {
		return err
	}
	err = d.Set("target_pipeline", s.Pipeline)
	if err != nil {
		return err
	}
	err = d.Set("pipeline_parameters", s.PipelineParameters)
	if err != nil {
		return err
	}
	return d.Set("wait_for_completion", s.WaitForCompletion)
}
