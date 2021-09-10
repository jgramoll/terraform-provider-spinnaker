package provider

import (
	"github.com/get-bridge/terraform-provider-spinnaker/client"
)

type stageEnabled struct {
	Expression string `mapstructure:"expression"`
	Type       string `mapstructure:"type"`
}

func fromClientStageEnabled(clientStageEnabled *client.StageEnabled) *[]*stageEnabled {
	if clientStageEnabled == nil {
		return nil
	}
	newStageEnabled := stageEnabled(*clientStageEnabled)
	newStageEnabledArray := []*stageEnabled{&newStageEnabled}
	return &newStageEnabledArray
}

func toClientStageEnabled(s *[]*stageEnabled) *client.StageEnabled {
	if s == nil || len(*s) == 0 {
		return nil
	}

	newStageEnabled := client.StageEnabled(*(*s)[0])
	return &newStageEnabled
}
