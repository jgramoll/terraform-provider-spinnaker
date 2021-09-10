package provider

import (
	"github.com/get-bridge/terraform-provider-spinnaker/client"
	"github.com/hashicorp/terraform/helper/schema"
)

type defaultNotification struct {
	ID      string      `mapstructure:"-"`
	Address string      `mapstructure:"address"`
	Message *[]*message `mapstructure:"message"`
	Type    string      `mapstructure:"type"`
	When    *[]*when    `mapstructure:"when"`
}

func newDefaultNotification() *defaultNotification {
	return &defaultNotification{}
}

func (n *defaultNotification) toClientNotification(level client.NotificationLevel) (*client.Notification, error) {
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

func (n *defaultNotification) fromClientNotification(cn *client.Notification) notification {
	return &defaultNotification{
		ID:      cn.ID,
		Address: cn.Address,
		Message: &[]*message{fromClientMessage(cn.Message)},
		Type:    cn.Type,
		When:    &[]*when{fromClientWhen(cn)},
	}
}

func (n *defaultNotification) setNotificationResourceData(d *schema.ResourceData) error {
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
