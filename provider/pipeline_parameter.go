package provider

import (
	"github.com/get-bridge/terraform-provider-spinnaker/client"
	"github.com/hashicorp/terraform/helper/schema"
)

type pipelineParameter struct {
	ID          string                      `mapstructure:"id"`
	Default     string                      `mapstructure:"default"`
	Description string                      `mapstructure:"description"`
	Label       string                      `mapstructure:"label"`
	Name        string                      `mapstructure:"name"`
	Options     *[]*pipelineParameterOption `mapstructure:"option"`
	Required    bool                        `mapstructure:"required"`
}

func fromClientPipelineParameter(pc *client.PipelineParameter) *pipelineParameter {
	return &pipelineParameter{
		ID:          pc.ID,
		Name:        pc.Name,
		Default:     pc.Default,
		Description: pc.Description,
		Label:       pc.Label,
		Options:     fromClientPipelineParameterOptions(pc.Options),
		Required:    pc.Required,
	}
}

func toClientPipelineParameter(p *pipelineParameter) *client.PipelineParameter {
	return &client.PipelineParameter{
		ID:          p.ID,
		Name:        p.Name,
		Default:     p.Default,
		Description: p.Description,
		HasOptions:  p.Options != nil && len(*p.Options) > 0,
		Label:       p.Label,
		Options:     toClientPipelineParameterOptions(p.Options),
		Required:    p.Required,
	}
}

func (parameter *pipelineParameter) setResourceData(d *schema.ResourceData) error {
	err := d.Set("default", parameter.Default)
	if err != nil {
		return err
	}
	err = d.Set("description", parameter.Description)
	if err != nil {
		return err
	}
	err = d.Set("label", parameter.Label)
	if err != nil {
		return err
	}
	err = d.Set("name", parameter.Name)
	if err != nil {
		return err
	}
	err = d.Set("option", parameter.Options)
	if err != nil {
		return err
	}
	return d.Set("required", parameter.Required)
}
