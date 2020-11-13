package provider

import (
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

type message struct {
	Complete string `mapstructure:"complete"`
	Failed   string `mapstructure:"failed"`
	Starting string `mapstructure:"starting"`
}

func toClientMessage(level client.NotificationLevel, m *[]*message) (client.Message, error) {
	if m == nil || len(*m) == 0 {
		return nil, nil
	}
	newMessage, err := client.NewMessage(level)
	if err != nil {
		return nil, err
	}
	message := (*m)[0]
	if message == nil {
		return nil, nil
	}

	if message.Complete != "" {
		newMessage.SetCompleteText(message.Complete)
	}
	if message.Failed != "" {
		newMessage.SetFailedText(message.Failed)
	}
	if message.Starting != "" {
		newMessage.SetStartingText(message.Starting)
	}
	return newMessage, nil
}

func fromClientMessage(cm client.Message) *message {
	if cm == nil {
		return nil
	}

	newMessage := message{}
	if cm.CompleteText() != "" {
		newMessage.Complete = cm.CompleteText()
	}
	if cm.FailedText() != "" {
		newMessage.Failed = cm.FailedText()
	}
	if cm.StartingText() != "" {
		newMessage.Starting = cm.StartingText()
	}
	return &newMessage
}
