package provider

import (
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

type message struct {
	Complete string
	Failed   string
	Starting string
}

func toClientMessage(level client.NotificationLevel, m *message) (client.Message, error) {
	newMessage, err := client.NewMessage(level)
	if err != nil {
		return nil, err
	}

	if m.Complete != "" {
		newMessage.SetCompleteText(m.Complete)
	}
	if m.Failed != "" {
		newMessage.SetFailedText(m.Failed)
	}
	if m.Starting != "" {
		newMessage.SetStartingText(m.Starting)
	}
	return newMessage, nil
}

func (m *message) fromClientMessage(cm client.Message) *message {
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
