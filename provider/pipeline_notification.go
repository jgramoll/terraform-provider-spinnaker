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

// TODO do we need this?
func (n notification) toClientNotification() client.Notification {
	return client.Notification{
		ID:      n.ID,
		Address: n.Address,
		Level:   n.Level,
		Message: n.Message.toClientMessage(),
		Type:    n.Type,
		When:    n.When.toClientWhen(),
	}
}

// TODO do we need this?
func (m message) toClientMessage() client.Message {
	return client.Message{
		Complete: client.MessageText{Text: m.Complete},
		Failed:   client.MessageText{Text: m.Failed},
		Starting: client.MessageText{Text: m.Starting},
	}
}

// TODO do we need this?
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

// TODO update to (pipeline *Pipeline)
func getNotification(notifications []client.Notification, notificationID string) (*client.Notification, error) {
	for _, notification := range notifications {
		if notification.ID == notificationID {
			return &notification, nil
		}
	}
	return nil, ErrNotificationNotFound
}

// TODO update to (pipeline *Pipeline)
func updateNotifications(pipeline *client.Pipeline, notification client.Notification) error {
	for i, t := range pipeline.Notifications {
		if t.ID == notification.ID {
			pipeline.Notifications[i] = notification
			return nil
		}
	}
	return ErrNotificationNotFound
}

// TODO update to (pipeline *Pipeline)
func deleteNotification(pipeline *client.Pipeline, notification client.Notification) error {
	for i, t := range pipeline.Notifications {
		if t.ID == notification.ID {
			pipeline.Notifications = append(pipeline.Notifications[:i], pipeline.Notifications[i+1:]...)
			return nil
		}
	}
	return ErrNotificationNotFound
}
