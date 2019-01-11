package client

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// ErrPipelineNotificationNotFound notification not found
var ErrPipelineNotificationNotFound = errors.New("notification not found")

// SerializableNotification for pipeline
type SerializableNotification struct {
	ID      *string           `json:"id"`
	Address string            `json:"address"`
	Level   NotificationLevel `json:"level"`
	Type    string            `json:"type"`
	When    []string          `json:"when"`
}

// Notification for pipeline
type Notification struct {
	SerializableNotification
	Message Message `json:"message"`
}

//GetNotification Get notification by id
func (pipeline *Pipeline) GetNotification(notificationID string) (*Notification, error) {
	notifications := *pipeline.Notifications
	for _, notification := range notifications {
		if notification.ID != nil && *notification.ID == notificationID {
			return notification, nil
		}
	}
	return nil, ErrPipelineNotificationNotFound
}

// UpdateNotification update notification
func (pipeline *Pipeline) UpdateNotification(notification *Notification) error {
	notifications := *pipeline.Notifications
	for i, t := range notifications {
		if *t.ID == *notification.ID {
			notifications[i] = notification
			return nil
		}
	}
	return ErrPipelineNotificationNotFound
}

//DeleteNotification delete notification
func (pipeline *Pipeline) DeleteNotification(notificationID string) error {
	notifications := *pipeline.Notifications
	for i, t := range notifications {
		if t.ID != nil && *t.ID == notificationID {
			notifications = append(notifications[:i], notifications[i+1:]...)
			pipeline.Notifications = &notifications
			return nil
		}
	}
	return ErrPipelineNotificationNotFound
}

func parseNotifications(notificationsHashInterface interface{}) (*[]*Notification, error) {
	notifications := []*Notification{}
	if notificationsHashInterface == nil {
		return &notifications, nil
	}

	notificationsToParse := notificationsHashInterface.([]interface{})
	for _, notificationInterface := range notificationsToParse {
		notificationMap := notificationInterface.(map[string]interface{})
		serializableNotification := SerializableNotification{}
		if err := mapstructure.Decode(notificationMap, &serializableNotification); err != nil {
			return nil, err
		}

		notification := Notification{SerializableNotification: serializableNotification}
		messageMap, ok := notificationMap["message"]
		if ok {
			message, err := parseMessage(serializableNotification.Level, messageMap.(map[string]interface{}))
			if err != nil {
				return nil, err
			}
			notification.Message = message
		}

		notifications = append(notifications, &notification)
	}
	return &notifications, nil
}
