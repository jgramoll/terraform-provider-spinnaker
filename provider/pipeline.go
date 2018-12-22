package provider

import (
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

// Pipeline deploy pipeline in application
type Pipeline struct {
	Application          string `mapstructure:"application"`
	Disabled             bool   `mapstructure:"disabled"`
	ID                   string `mapstructure:"id"`
	KeepWaitingPipelines bool   `mapstructure:"keep_waiting_pipelines"`
	LimitConcurrent      bool   `mapstructure:"limit_concerrent"`
	Name                 string `mapstructure:"name"`
	Index                int    `mapstructure:"index"`
}

// ToClientPipeline convert to client pipeline
// TODO better way?
func (pipeline *Pipeline) ToClientPipeline() *client.Pipeline {
	return &client.Pipeline{
		PipelineWithoutStages: client.PipelineWithoutStages{
			Application:          pipeline.Application,
			Disabled:             pipeline.Disabled,
			ID:                   pipeline.ID,
			KeepWaitingPipelines: pipeline.KeepWaitingPipelines,
			LimitConcurrent:      pipeline.LimitConcurrent,
			Name:                 pipeline.Name,
			Index:                pipeline.Index,
		},
	}
}
