package client

import (
	"github.com/mitchellh/mapstructure"
)

var messageFactories = map[NotificationLevel]func() interface{}{}

// MessageText for Pipeline Notification
type MessageText struct {
	Text string `json:"text"`
}

// Message for Notification
type Message interface {
	SetCompleteText(string)
	CompleteText() string

	SetFailedText(string)
	FailedText() string

	SetStartingText(string)
	StartingText() string
}

// NewMessage new message
func NewMessage(level NotificationLevel) Message {
	return messageFactories[level]().(Message)
}

func parseMessage(level NotificationLevel, messageMap map[string]interface{}) (Message, error) {
	factory := messageFactories[level]
	if factory == nil {
		return nil, ErrInvalidNotificatoinLevel
	}
	message := factory()

	if err := mapstructure.Decode(messageMap, message); err != nil {
		return nil, err
	}
	return message.(Message), nil
}
