package provider

import (
	"github.com/get-bridge/terraform-provider-spinnaker/client"
)

type manualJudgementNotificationMessage struct {
	ManualJudgmentContinue string `mapstructure:"manual_judgment_continue"`
	ManualJudgmentStop     string `mapstructure:"manual_judgment_stop"`
}

func toClientManualJudgementNotificationMessage(level client.NotificationLevel, m *[]*manualJudgementNotificationMessage) (client.Message, error) {
	if m == nil || len(*m) == 0 {
		return nil, nil
	}
	message := (*m)[0]
	if message == nil {
		return nil, nil
	}

	newMessage, err := client.NewMessage(level)
	if err != nil {
		return nil, err
	}

	if message.ManualJudgmentContinue != "" {
		newMessage.SetManualJudgmentContinueText(message.ManualJudgmentContinue)
	}
	if message.ManualJudgmentStop != "" {
		newMessage.SetManualJudgmentStopText(message.ManualJudgmentStop)
	}
	return newMessage, nil
}

func fromClientManualJudgementNotificationMessage(cm client.Message) *manualJudgementNotificationMessage {
	if cm == nil {
		return nil
	}

	newMessage := manualJudgementNotificationMessage{}
	if cm.ManualJudgmentContinueText() != "" {
		newMessage.ManualJudgmentContinue = cm.ManualJudgmentContinueText()
	}
	if cm.ManualJudgmentStopText() != "" {
		newMessage.ManualJudgmentStop = cm.ManualJudgmentStopText()
	}
	return &newMessage
}
