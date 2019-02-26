package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

type pipelineParameter struct {
	Default     string                      `mapstructure:"default"`
	Description string                      `mapstructure:"description"`
	Label       string                      `mapstructure:"label"`
	Name        string                      `mapstructure:"name"`
	Options     *[]*pipelineParameterOption `mapstructure:"options"`
	Required    bool                        `mapstructure:"required"`
}

func toClientPipelineConfig(parameters *[]*pipelineParameter) *[]*client.PipelineParameter {
	if parameters == nil {
		return nil
	}

	config := []*client.PipelineParameter{}
	for _, pc := range *parameters {
		config = append(config, &client.PipelineParameter{
			Name:        pc.Name,
			Default:     pc.Default,
			Description: pc.Description,
			HasOptions:  pc.Options != nil && len(*pc.Options) > 0,
			Label:       pc.Label,
			Options:     toClientPipelineParameterOptions(pc.Options),
			Required:    pc.Required,
		})
	}

	return &config
}

func fromClientPipelineConfig(parameters *[]*client.PipelineParameter) *[]*pipelineParameter {
	if parameters == nil {
		return nil
	}

	config := []*pipelineParameter{}
	for _, pc := range *parameters {
		config = append(config, &pipelineParameter{
			Name:        pc.Name,
			Default:     pc.Default,
			Description: pc.Description,
			Label:       pc.Label,
			Options:     fromClientPipelineParameterOptions(pc.Options),
			Required:    pc.Required,
		})
	}

	return &config
}

func pipelineParametersFromResourceData(d *schema.ResourceData) *[]*client.PipelineParameter {
	parameters := []*client.PipelineParameter{}
	state := d.Get("parameter").([]interface{})
	for _, paramInterface := range state {
		param := paramInterface.(map[string]interface{})
		options := pipelineParameterOptionFromMap(param["option"].([]interface{}))
		parameters = append(parameters, &client.PipelineParameter{
			Default:     param["default"].(string),
			Description: param["description"].(string),
			HasOptions:  len(*options) > 0,
			Label:       param["label"].(string),
			Name:        param["name"].(string),
			Options:     options,
			Required:    param["required"].(bool),
		})
	}

	return &parameters
}
