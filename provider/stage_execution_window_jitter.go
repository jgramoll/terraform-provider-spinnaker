package provider

import (
	"github.com/get-bridge/terraform-provider-spinnaker/client"
)

type stageExecutionWindowJitter struct {
	Enabled    bool `mapstructure:"enabled"`
	MaxDelay   int  `mapstructure:"max_delay"`
	MinDelay   int  `mapstructure:"min_delay"`
	SkipManual bool `mapstructure:"skip_manual"`
}

func toClientWindowJitter(jitter *[]*stageExecutionWindowJitter) *client.StageExecutionWindowJitter {
	if jitter == nil || len(*jitter) == 0 {
		return nil
	}
	newJitter := client.StageExecutionWindowJitter(*(*jitter)[0])
	return &newJitter
}

func fromClientExecutionJitter(clientWindowJitter *client.StageExecutionWindowJitter) *[]*stageExecutionWindowJitter {
	if clientWindowJitter == nil {
		return nil
	}
	newJitter := stageExecutionWindowJitter(*clientWindowJitter)
	return &[]*stageExecutionWindowJitter{&newJitter}
}
