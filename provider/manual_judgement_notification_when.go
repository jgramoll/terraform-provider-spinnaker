package provider

import (
	"github.com/get-bridge/terraform-provider-spinnaker/client"
)

type manualJudgementNotificationWhen struct {
	ManualJudgment         bool `mapstructure:"manual_judgment"`
	ManualJudgmentContinue bool `mapstructure:"manual_judgment_continue"`
	ManualJudgmentStop     bool `mapstructure:"manual_judgment_stop"`
}

func newManualJudgementNotificationWhen() *manualJudgementNotificationWhen {
	return &manualJudgementNotificationWhen{}
}

func toClientManualJudgementNotificationWhen(level client.NotificationLevel, w *manualJudgementNotificationWhen) *[]string {
	clientWhen := []string{}
	if w == nil {
		return &clientWhen
	}
	if level == client.NotificationLevelPipeline {
		panic("NOT IMPLEMENTED")
	} else if level == client.NotificationLevelStage {
		if w.ManualJudgment {
			clientWhen = append(clientWhen, client.ManualJudgmentKey)
		}
		if w.ManualJudgmentContinue {
			clientWhen = append(clientWhen, client.ManualJudgmentContinueKey)
		}
		if w.ManualJudgmentStop {
			clientWhen = append(clientWhen, client.ManualJudgmentStopKey)
		}
	}
	return &clientWhen
}

func fromClientManualJudgementNotificationWhen(cn *client.Notification) *manualJudgementNotificationWhen {
	w := newManualJudgementNotificationWhen()
	for _, cw := range cn.When {
		switch cw {
		case client.ManualJudgmentKey:
			w.ManualJudgment = true
		case client.ManualJudgmentContinueKey:
			w.ManualJudgmentContinue = true
		case client.ManualJudgmentStopKey:
			w.ManualJudgmentStop = true
		}
	}
	return w
}
