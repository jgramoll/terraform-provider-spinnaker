package provider

import (
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

type stageExecutionWindowJitter struct {
	Enabled    bool `mapstructure:"enabled"`
	MaxDelay   int  `mapstructure:"max_delay"`
	MinDelay   int  `mapstructure:"min_delay"`
	SkipManual bool `mapstructure:"skip_manual"`
}

func toClientWindowJitter(jitter *stageExecutionWindowJitter) *client.StageExecutionWindowJitter {
	if jitter == nil {
		return nil
	}
	newJitter := client.StageExecutionWindowJitter(*jitter)
	return &newJitter
}

func fromClientExecutionJitter(clientWindowJitter *client.StageExecutionWindowJitter) *stageExecutionWindowJitter {
	if clientWindowJitter == nil {
		return nil
	}
	newJitter := stageExecutionWindowJitter(*clientWindowJitter)
	return &newJitter
}
