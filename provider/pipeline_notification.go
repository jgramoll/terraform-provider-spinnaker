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
	ID      string
	Address string
	Level   string
	Message message
	Type    string
	When    when
}

func (n notification) toClientNotification() *client.Notification {
	return &client.Notification{
		ID:      n.ID,
		Address: n.Address,
		Level:   n.Level,
		Message: n.Message.toClientMessage(),
		Type:    n.Type,
		When:    n.When.toClientWhen(),
	}
}

func (m message) toClientMessage() client.Message {
	return client.Message{
		Complete: client.MessageText{Text: m.Complete},
		Failed:   client.MessageText{Text: m.Failed},
		Starting: client.MessageText{Text: m.Starting},
	}
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
