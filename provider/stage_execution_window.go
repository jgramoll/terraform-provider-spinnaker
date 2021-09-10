package provider

import (
	"github.com/get-bridge/terraform-provider-spinnaker/client"
)

type stageExecutionWindow struct {
	Days      []int                             `mapstructure:"days"`
	Jitter    *[]*stageExecutionWindowJitter    `mapstructure:"jitter"`
	Whitelist *[]*stageExecutionWindowWhitelist `mapstructure:"whitelist"`
}

func toClientExecutionWindow(clientWindow *[]*stageExecutionWindow) *client.StageExecutionWindow {
	if clientWindow == nil || len(*clientWindow) == 0 {
		return nil
	}
	newWindow := client.StageExecutionWindow{}
	for _, w := range *clientWindow {
		newWindow.Days = w.Days
		newWindow.Whitelist = toClientWindowWhitelist(w.Whitelist)
		newWindow.Jitter = toClientWindowJitter(w.Jitter)
		break
	}
	return &newWindow
}

func fromClientExecutionWindow(clientWindow *client.StageExecutionWindow) *[]*stageExecutionWindow {
	if clientWindow == nil {
		return nil
	}
	newWindow := stageExecutionWindow{
		Days:      clientWindow.Days,
		Whitelist: fromClientExecutionWhitelist(clientWindow.Whitelist),
		Jitter:    fromClientExecutionJitter(clientWindow.Jitter),
	}
	newWindowList := []*stageExecutionWindow{&newWindow}
	return &newWindowList
}
