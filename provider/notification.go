package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

type notification struct {
	ID      string      `mapstructure:"-"`
	Address string      `mapstructure:"address"`
	Message *[]*message `mapstructure:"message"`
	Type    string      `mapstructure:"type"`
	When    *[]*when    `mapstructure:"when"`
}

func (n *notification) toClientNotification(level client.NotificationLevel) (*client.Notification, error) {
	message, err := toClientMessage(level, n.Message)
	if err != nil {
		return nil, err
	}
	return &client.Notification{
		ID:      n.ID,
		Address: n.Address,
		Level:   level,
		Type:    n.Type,
		When:    *toClientWhen(level, (*n.When)[0]),
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
		newNotifications = append(newNotifications, fromClientNotification(cn))
	}
	return &newNotifications
}

func fromClientNotification(cn *client.Notification) *notification {
	return &notification{
		ID:      cn.ID,
		Address: cn.Address,
		Message: &[]*message{fromClientMessage(cn.Message)},
		Type:    cn.Type,
		When:    &[]*when{fromClientWhen(cn)},
	}
}

func (n *notification) setNotificationResourceData(d *schema.ResourceData) error {
	err := d.Set("address", n.Address)
	if err != nil {
		return err
	}
	err = d.Set("message", n.Message)
	if err != nil {
		return err
	}
	err = d.Set("type", n.Type)
	if err != nil {
		return err
	}
	return d.Set("when", n.When)
}
