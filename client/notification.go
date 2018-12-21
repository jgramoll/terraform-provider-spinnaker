package client

import (
	"errors"
)

// PipelineCompleteKey for pipeline complete
const PipelineCompleteKey = "pipeline.complete"

// PipelineFailedKey for pipeline failed
const PipelineFailedKey = "pipeline.failed"

// PipelineStartingKey for pipeline starting
const PipelineStartingKey = "pipeline.starting"

// ErrPipelineNotificationNotFound notification not found
var ErrPipelineNotificationNotFound = errors.New("notification not found")

// MessageText for Pipeline Notification
type MessageText struct {
	Text string `json:"text"`
}

// Message for Pipeline Notification
type Message struct {
	Complete MessageText `json:"pipeline.complete" mapstructure:"pipeline.complete"`
	Failed   MessageText `json:"pipeline.failed" mapstructure:"pipeline.failed"`
	Starting MessageText `json:"pipeline.starting" mapstructure:"pipeline.starting"`
}

// Notification for Pipeline
type Notification struct {
	ID      string   `json:"id"`
	Address string   `json:"address"`
	Level   string   `json:"level"`
	Message Message  `json:"message"`
	Type    string   `json:"type"`
	When    []string `json:"when"`
}

//GetNotification Get notification by id
func (pipeline *Pipeline) GetNotification(notificationID string) (*Notification, error) {
	for _, notification := range pipeline.Notifications {
		if notification.ID == notificationID {
			return &notification, nil
		}
	}
	return nil, ErrPipelineNotificationNotFound
}

// UpdateNotification update notification
func (pipeline *Pipeline) UpdateNotification(notification *Notification) error {
	for i, t := range pipeline.Notifications {
		if t.ID == notification.ID {
			pipeline.Notifications[i] = *notification
			return nil
		}
	}
	return ErrPipelineNotificationNotFound
}

//DeleteNotification delete notification
func (pipeline *Pipeline) DeleteNotification(notificationID string) error {
	for i, t := range pipeline.Notifications {
		if t.ID == notificationID {
			pipeline.Notifications = append(pipeline.Notifications[:i], pipeline.Notifications[i+1:]...)
			return nil
		}
	}
	return ErrPipelineNotificationNotFound
}
