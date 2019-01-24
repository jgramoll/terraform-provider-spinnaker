package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

type pipelineParameterOption struct {
	Value string `mapstructure:"value"`
}

type pipelineParameter struct {
	Description string                    `mapstructure:"description"`
	Name        string                    `mapstructure:"name"`
	Options     []pipelineParameterOption `mapstructure:"option"`
	Required    bool                      `mapstructure:"required"`
}

func PipelineParametersFromResourceData(d *schema.ResourceData) *[]*client.PipelineParameter {
	parameters := []*client.PipelineParameter{}
	state := d.Get("parameter").([]interface{})
	for _, paramInterface := range state {
		param := paramInterface.(map[string]interface{})
		options := PipelineParameterOptionFromMap(param["option"].([]interface{}))
		parameters = append(parameters, &client.PipelineParameter{
			Name:        param["name"].(string),
			Description: param["description"].(string),
			HasOptions:  len(*options) > 0,
			Options:     options,
			Required:    param["required"].(bool),
		})
	}

	return &parameters
}

func PipelineParameterOptionFromMap(options []interface{}) *[]client.PipelineParameterOption {
	clientOptions := []client.PipelineParameterOption{}
	for _, optionInterface := range options {
		option := optionInterface.(map[string]interface{})
		clientOptions = append(clientOptions, client.PipelineParameterOption{
			Value: option["value"].(string),
		})
	}

	return &clientOptions
}
