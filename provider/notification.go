package provider

import (
	"github.com/get-bridge/terraform-provider-spinnaker/client"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/mitchellh/mapstructure"
)

type notification interface {
	toClientNotification(level client.NotificationLevel) (*client.Notification, error)
	fromClientNotification(cn *client.Notification) notification
	setNotificationResourceData(d *schema.ResourceData) error
}

func toClientNotifications(notificationFactory func() notification, notifications *[]map[string]interface{}) (*[]*client.Notification, error) {
	clientNotifications := []*client.Notification{}
	if notifications != nil {
		for _, notificationMap := range *notifications {
			n := notificationFactory()
			if err := mapstructure.Decode(notificationMap, n); err != nil {
				return nil, err
			}
			cn, err := n.toClientNotification(client.NotificationLevelStage)
			if err != nil {
				return nil, err
			}
			clientNotifications = append(clientNotifications, cn)
		}
	}
	return &clientNotifications, nil
}

func fromClientNotifications(notificationFactory func() notification, notifications *[]*client.Notification) (*[]map[string]interface{}, error) {
	if notifications == nil {
		return nil, nil
	}

	newNotifications := []map[string]interface{}{}
	for _, cn := range *notifications {
		notificationMap := map[string]interface{}{}
		if err := mapstructure.Decode(notificationFactory().fromClientNotification(cn), &notificationMap); err != nil {
			return nil, err
		}
		newNotifications = append(newNotifications, notificationMap)
	}
	return &newNotifications, nil
}
