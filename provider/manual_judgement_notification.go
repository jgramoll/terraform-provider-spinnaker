package provider

import (
	"github.com/get-bridge/terraform-provider-spinnaker/client"
	"github.com/hashicorp/terraform/helper/schema"
)

type manualJudgementNotification struct {
	ID      string `mapstructure:"-"`
	Address string `mapstructure:"address"`
	Type    string `mapstructure:"type"`

	Message *[]*manualJudgementNotificationMessage `mapstructure:"message"`
	When    *[]*manualJudgementNotificationWhen    `mapstructure:"when"`
}

func newManualJudgementNotification() *manualJudgementNotification {
	return &manualJudgementNotification{}
}

func newManualJudgementNotificationInterface() notification {
	return newManualJudgementNotification()
}

func (n *manualJudgementNotification) toClientNotification(level client.NotificationLevel) (*client.Notification, error) {
	message, err := toClientManualJudgementNotificationMessage(level, n.Message)
	if err != nil {
		return nil, err
	}
	return &client.Notification{
		ID:      n.ID,
		Address: n.Address,
		Level:   level,
		Type:    n.Type,
		When:    *toClientManualJudgementNotificationWhen(level, (*n.When)[0]),
		Message: message,
	}, nil
}

func (n *manualJudgementNotification) fromClientNotification(cn *client.Notification) notification {
	return &manualJudgementNotification{
		ID:      cn.ID,
		Address: cn.Address,
		Message: &[]*manualJudgementNotificationMessage{fromClientManualJudgementNotificationMessage(cn.Message)},
		Type:    cn.Type,
		When:    &[]*manualJudgementNotificationWhen{fromClientManualJudgementNotificationWhen(cn)},
	}
}

func (n *manualJudgementNotification) setNotificationResourceData(d *schema.ResourceData) error {
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
