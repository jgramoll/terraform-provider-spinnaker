package client

import (
	"errors"
)

// ErrInvalidNotificationLevel invalid notification level
var ErrInvalidNotificationLevel = errors.New("Invalid notification level")

// NotificationLevel level of notification
type NotificationLevel string

// NotificationLevelStage stage level notifaction
const NotificationLevelStage NotificationLevel = "stage"

// NotificationLevelPipeline pipeline level notifaction
const NotificationLevelPipeline NotificationLevel = "pipeline"
