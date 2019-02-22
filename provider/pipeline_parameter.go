package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

type pipelineParameterOption struct {
	Value string `mapstructure:"value"`
}

type pipelineParameter struct {
	Default     string                     `mapstructure:"default"`
	Description string                     `mapstructure:"description"`
	Label       string                     `mapstructure:"label"`
	Name        string                     `mapstructure:"name"`
	Options     []*pipelineParameterOption `mapstructure:"options"`
	Required    bool                       `mapstructure:"required"`
}

func (parameters *pipelineParameter) ToClientPipelineParameterOption() *[]*client.PipelineParameterOption {
	options := []*client.PipelineParameterOption{}

	for _, option := range parameters.Options {
		options = append(options, &client.PipelineParameterOption{
			Value: option.Value,
		})
	}

	return &options
}

func pipelineParametersFromResourceData(d *schema.ResourceData) *[]*client.PipelineParameter {
	parameters := []*client.PipelineParameter{}
	state := d.Get("parameter").([]interface{})
	for _, paramInterface := range state {
		param := paramInterface.(map[string]interface{})
		options := PipelineParameterOptionFromMap(param["option"].([]interface{}))
		parameters = append(parameters, &client.PipelineParameter{
			Default:     param["default"].(string),
			Description: param["description"].(string),
			HasOptions:  len(*options) > 0,
			Label:       param["label"].(string),
			Name:        param["name"].(string),
			Options:     *options,
			Required:    param["required"].(bool),
		})
	}

	return &parameters
}

func PipelineParameterOptionFromMap(options []interface{}) *[]*client.PipelineParameterOption {
	clientOptions := []*client.PipelineParameterOption{}
	for _, optionInterface := range options {
		option := optionInterface.(map[string]interface{})
		clientOptions = append(clientOptions, &client.PipelineParameterOption{
			Value: option["value"].(string),
		})
	}

	return &clientOptions
}
