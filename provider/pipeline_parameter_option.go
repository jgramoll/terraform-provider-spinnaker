package provider

import (
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

type pipelineParameterOption struct {
	Value string `mapstructure:"value"`
}

func toClientPipelineParameterOptions(parameters *[]*pipelineParameterOption) *[]*client.PipelineParameterOption {
	if parameters == nil {
		return nil
	}
	options := []*client.PipelineParameterOption{}
	for _, option := range *parameters {
		options = append(options, &client.PipelineParameterOption{
			Value: option.Value,
		})
	}

	return &options
}

func fromClientPipelineParameterOptions(parameters *[]*client.PipelineParameterOption) *[]*pipelineParameterOption {
	if parameters == nil {
		return nil
	}
	options := []*pipelineParameterOption{}
	for _, option := range *parameters {
		options = append(options, &pipelineParameterOption{
			Value: option.Value,
		})
	}

	return &options
}

func pipelineParameterOptionFromMap(options []interface{}) *[]*client.PipelineParameterOption {
	clientOptions := []*client.PipelineParameterOption{}
	for _, optionInterface := range options {
		option := optionInterface.(map[string]interface{})
		clientOptions = append(clientOptions, &client.PipelineParameterOption{
			Value: option["value"].(string),
		})
	}

	return &clientOptions
}
