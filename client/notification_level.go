package client

import (
	"errors"
)

// ErrInvalidNotificatoinLevel invalid notification level
var ErrInvalidNotificatoinLevel = errors.New("Invalid notification level")

// NotificationLevel level of notification
type NotificationLevel string

// NotificationLevelStage stage level notifaction
const NotificationLevelStage NotificationLevel = "stage"

// NotificationLevelPipeline pipeline level notifaction
const NotificationLevelPipeline NotificationLevel = "pipeline"
