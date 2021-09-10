package provider

import (
	"github.com/get-bridge/terraform-provider-spinnaker/client"
)

type canaryConfigJudges []*canaryConfigJudge

type canaryConfigJudge struct {
	Name string `mapstructure:"name"`
	// JudgeConfigurations map[string]interface{} `mapstructure:"judge_configurations"`
}

func (judges *canaryConfigJudges) toClientJudge() *client.CanaryConfigJudge {
	for _, judge := range *judges {
		return &client.CanaryConfigJudge{
			Name: judge.Name,
			// JudgeConfigurations: judge.JudgeConfigurations,
		}
	}
	return nil
}

func (*canaryConfigJudges) fromClientJudge(judge *client.CanaryConfigJudge) *canaryConfigJudges {
	return &canaryConfigJudges{&canaryConfigJudge{
		Name: judge.Name,
	}}
}
