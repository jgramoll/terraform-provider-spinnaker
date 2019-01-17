package provider

import (
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

type notification struct {
	ID      string
	Address string
	Message *[]*message
	Type    string
	When    *[]*when
}

func (n *notification) toClientNotification(level client.NotificationLevel) (*client.Notification, error) {
	message, err := toClientMessage(level, (*n.Message)[0])
	if err != nil {
		return nil, err
	}
	return &client.Notification{
		SerializableNotification: client.SerializableNotification{
			ID:      n.ID,
			Address: n.Address,
			Level:   level,
			Type:    n.Type,
			When:    *toClientWhen(level, (*n.When)[0]),
		},
		Message: message,
	}, nil
}

func toClientNotifications(notifications *[]*notification) (*[]*client.Notification, error) {
	clientNotifications := []*client.Notification{}
	if notifications != nil {
		for _, n := range *notifications {
			cn, err := n.toClientNotification(client.NotificationLevelStage)
			if err != nil {
				return nil, err
			}
			clientNotifications = append(clientNotifications, cn)
		}
	}
	return &clientNotifications, nil
}

func fromClientNotifications(notifications *[]*client.Notification) *[]*notification {
	if notifications == nil {
		return nil
	}

	newNotifications := []*notification{}
	for _, cn := range *notifications {
		newNotifications = append(newNotifications, &notification{
			ID:      cn.ID,
			Address: cn.Address,
			Message: &[]*message{(&message{}).fromClientMessage(cn.Message)},
			Type:    cn.Type,
			When:    &[]*when{(&when{}).fromClientWhen(cn)},
		})
	}
	return &newNotifications
}
