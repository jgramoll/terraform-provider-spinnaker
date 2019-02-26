package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
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
		Name:        p.Name,
		Default:     p.Default,
		Description: p.Description,
		HasOptions:  p.Options != nil && len(*p.Options) > 0,
		Label:       p.Label,
		Options:     toClientPipelineParameterOptions(p.Options),
		Required:    p.Required,
	}
}

func toClientPipelineConfig(parameters *[]*pipelineParameter) *[]*client.PipelineParameter {
	if parameters == nil {
		return nil
	}

	config := []*client.PipelineParameter{}
	for _, p := range *parameters {
		config = append(config, toClientPipelineParameter(p))
	}

	return &config
}

func fromClientPipelineConfig(parameters *[]*client.PipelineParameter) *[]*pipelineParameter {
	if parameters == nil {
		return nil
	}

	config := []*pipelineParameter{}
	for _, pc := range *parameters {
		config = append(config, fromClientPipelineParameter(pc))
	}

	return &config
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
