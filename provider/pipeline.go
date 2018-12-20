package provider

import (
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

// Pipeline deploy pipeline in application
type Pipeline struct {
	Application string
	ID          string
	Name        string
	Index       int
}

// ToClientPipeline convert to client pipeline
// TODO better way?
func (pipeline *Pipeline) ToClientPipeline() *client.Pipeline {
	return &client.Pipeline{
		PipelineWithoutStages: client.PipelineWithoutStages{
			Application: pipeline.Application,
			ID:          pipeline.ID,
			Name:        pipeline.Name,
			Index:       pipeline.Index,
		},
	}
}
