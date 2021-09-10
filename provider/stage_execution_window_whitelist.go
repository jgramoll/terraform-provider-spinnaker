package provider

import (
	"github.com/get-bridge/terraform-provider-spinnaker/client"
)

type stageExecutionWindowWhitelist struct {
	EndHour   int `mapstructure:"end_hour"`
	EndMin    int `mapstructure:"end_min"`
	StartHour int `mapstructure:"start_hour"`
	StartMin  int `mapstructure:"start_min"`
}

func toClientWindowWhitelist(w *[]*stageExecutionWindowWhitelist) *[]*client.StageExecutionWindowWhitelist {
	if w == nil || len(*w) == 0 {
		return nil
	}
	whitelists := []*client.StageExecutionWindowWhitelist{}
	for _, w := range *w {
		whitelist := client.StageExecutionWindowWhitelist(*w)
		whitelists = append(whitelists, &whitelist)
	}
	return &whitelists
}

func fromClientExecutionWhitelist(clientWindowWhitelist *[]*client.StageExecutionWindowWhitelist) *[]*stageExecutionWindowWhitelist {
	if clientWindowWhitelist == nil || len(*clientWindowWhitelist) == 0 {
		return nil
	}
	newWhitelists := []*stageExecutionWindowWhitelist{}
	for _, w := range *clientWindowWhitelist {
		whitelist := stageExecutionWindowWhitelist(*w)
		newWhitelists = append(newWhitelists, &whitelist)
	}
	return &newWhitelists
}
