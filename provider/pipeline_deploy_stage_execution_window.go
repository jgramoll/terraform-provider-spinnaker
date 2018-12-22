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

type stageExecutionWindowWhitelist struct {
	EndHour   int `mapstructure:"end_hour"`
	EndMin    int `mapstructure:"end_min"`
	StartHour int `mapstructure:"start_hour"`
	StartMin  int `mapstructure:"start_min"`
}

type stageExecutionWindow struct {
	Days      []int                           `mapstructure:"days"`
	Jitter    []stageExecutionWindowJitter    `mapstructure:"jitter"`
	Whitelist []stageExecutionWindowWhitelist `mapstructure:"whitelist"`
}

func (w *stageExecutionWindow) toClientWindowWhitelist() *[]client.StageExecutionWindowWhitelist {
	whitelists := []client.StageExecutionWindowWhitelist{}
	for _, whitelist := range w.Whitelist {
		whitelists = append(whitelists, client.StageExecutionWindowWhitelist(whitelist))
	}
	// TODO why does `[]client.StageExecutionWindowWhitelist(w.Whitelist)` not work
	return &whitelists
}

func (w *stageExecutionWindow) toClientExecutionWindow() *client.StageExecutionWindow {
	newWindow := client.StageExecutionWindow{
		Days:      w.Days,
		Whitelist: *w.toClientWindowWhitelist(),
	}
	if len(w.Jitter) > 0 {
		newJitter := client.StageExecutionWindowJitter(w.Jitter[0])
		newWindow.Jitter = &newJitter
	}

	return &newWindow
}

func (w *stageExecutionWindow) fromClientWindow(clientWindow *client.StageExecutionWindow) *stageExecutionWindow {
	newWindow := stageExecutionWindow{
		Days:      clientWindow.Days,
		Whitelist: []stageExecutionWindowWhitelist(w.Whitelist),
	}
	if clientWindow.Jitter != nil {
		newWindow.Jitter = []stageExecutionWindowJitter{stageExecutionWindowJitter(*clientWindow.Jitter)}
	}
	return &newWindow
}
