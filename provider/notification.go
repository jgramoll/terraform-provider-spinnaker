package provider

import (
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

type message struct {
	Complete string
	Failed   string
	Starting string
}

type when struct {
	Complete string
	Starting string
	Failed   string
}

type notification struct {
	ID      *string
	Address string
	Message []*message
	Type    string
	When    []*when
}

func (n notification) toClientNotification(level client.NotificationLevel) *client.Notification {
	return &client.Notification{
		SerializableNotification: client.SerializableNotification{
			ID:      n.ID,
			Address: n.Address,
			Level:   level,
			Type:    n.Type,
			When:    n.When[0].toClientWhen(),
		},
		Message: n.Message[0].toClientMessage(level),
	}
}

func (m message) toClientMessage(level client.NotificationLevel) client.Message {
	newMessage := client.NewMessage(level)

	if m.Complete != "" {
		newMessage.SetCompleteText(m.Complete)
	}
	if m.Failed != "" {
		newMessage.SetFailedText(m.Failed)
	}
	if m.Starting != "" {
		newMessage.SetStartingText(m.Starting)
	}
	return newMessage
}

func (w when) toClientWhen() []string {
	clientWhen := []string{}
	if w.Complete == "1" {
		clientWhen = append(clientWhen, client.PipelineCompleteKey)
	}
	if w.Failed == "1" {
		clientWhen = append(clientWhen, client.PipelineFailedKey)
	}
	if w.Starting == "1" {
		clientWhen = append(clientWhen, client.PipelineStartingKey)
	}
	return clientWhen
}

func (n notification) fromClientNotification(cn *client.Notification) *notification {
	return &notification{
		ID:      cn.ID,
		Address: cn.Address,
		Message: []*message{message{}.fromClientMessage(cn.Message)},
		Type:    cn.Type,
		When:    []*when{when{}.fromClientWhen(cn)},
	}
}

func (m message) fromClientMessage(cm client.Message) *message {
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

func (w when) fromClientWhen(cn *client.Notification) *when {
	newWhen := when{}
	for _, cw := range cn.When {
		switch cw {
		case client.StageCompleteKey:
			newWhen.Complete = "1"
		case client.PipelineCompleteKey:
			newWhen.Complete = "1"
		case client.StageFailedKey:
			newWhen.Failed = "1"
		case client.PipelineFailedKey:
			newWhen.Failed = "1"
		case client.StageStartingKey:
			newWhen.Starting = "1"
		case client.PipelineStartingKey:
			newWhen.Starting = "1"
		}
	}
	return &newWhen
}
