package client

import (
	"github.com/mitchellh/mapstructure"
)

var messageFactories = map[NotificationLevel]func() Message{}

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

	SetManualJudgmentContinueText(string)
	ManualJudgmentContinueText() string

	SetManualJudgmentStopText(string)
	ManualJudgmentStopText() string
}

// NewMessage new message
func NewMessage(level NotificationLevel) (Message, error) {
	factory := messageFactories[level]
	if factory == nil {
		return nil, ErrInvalidNotificationLevel
	}
	return factory(), nil
}

func parseMessage(level NotificationLevel, messageMap map[string]interface{}) (Message, error) {
	message, err := NewMessage(level)
	if err != nil {
		return nil, err
	}
	if err := mapstructure.Decode(messageMap, message); err != nil {
		return nil, err
	}
	return message, nil
}
