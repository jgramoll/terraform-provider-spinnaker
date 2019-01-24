package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

// Pipeline deploy pipeline in application
type Pipeline struct {
	Application          string               `mapstructure:"application"`
	Disabled             bool                 `mapstructure:"disabled"`
	ID                   string               `mapstructure:"id"`
	KeepWaitingPipelines bool                 `mapstructure:"keep_waiting_pipelines"`
	LimitConcurrent      bool                 `mapstructure:"limit_concerrent"`
	Name                 string               `mapstructure:"name"`
	Index                int                  `mapstructure:"index"`
	ParameterConfig      []*pipelineParameter `mapstructure:"parameter"`
}

// ToClientPipeline convert to client pipeline
// TODO better way?
func (pipeline *Pipeline) ToClientPipeline() *client.Pipeline {
	return &client.Pipeline{
		SerializablePipeline: client.SerializablePipeline{
			Application:          pipeline.Application,
			Disabled:             pipeline.Disabled,
			ID:                   pipeline.ID,
			KeepWaitingPipelines: pipeline.KeepWaitingPipelines,
			LimitConcurrent:      pipeline.LimitConcurrent,
			Name:                 pipeline.Name,
			Index:                pipeline.Index,
			ParameterConfig:      pipeline.ToClientPipelineConfig(),
		},
	}
}

func (pipeline *Pipeline) ToClientPipelineConfig() *[]*client.PipelineParameter {
	config := []*client.PipelineParameter{}

	for _, pc := range pipeline.ParameterConfig {
		config = append(config, &client.PipelineParameter{
			Name:        pc.Name,
			Description: pc.Description,
			HasOptions:  len(pc.Options) > 0,
			// Options:     &[]client.PipelineParameterOption(pc.Options),
			Required: pc.Required,
		})
	}

	return &config
}

func SetResourceData(pipeline *client.Pipeline, d *schema.ResourceData) {
	d.SetId(pipeline.ID)
	d.Set(ApplicationKey, pipeline.Application)
	d.Set("name", pipeline.Name)
	d.Set("index", pipeline.Index)
	d.Set("disabled", pipeline.Disabled)
	d.Set("keep_waiting_pipelines", pipeline.KeepWaitingPipelines)
	d.Set("limit_concurrent", pipeline.LimitConcurrent)
	d.Set("parameter", pipeline.ParameterConfig)
}

func PipelineFromResourceData(pipeline *client.Pipeline, d *schema.ResourceData) {
	pipeline.Index = d.Get("index").(int)
	pipeline.Application = d.Get(ApplicationKey).(string)
	pipeline.Name = d.Get("name").(string)
	pipeline.Disabled = d.Get("disabled").(bool)
	pipeline.KeepWaitingPipelines = d.Get("keep_waiting_pipelines").(bool)
	pipeline.LimitConcurrent = d.Get("limit_concurrent").(bool)
	pipeline.ParameterConfig = PipelineParametersFromResourceData(d)
}
