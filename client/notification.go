package client

// Notification for Pipeline
type Notification struct {
	ID      string   `json:"id"`
	Address string   `json:"address"`
	Level   string   `json:"level"`
	Message Message  `json:"message"`
	Type    string   `json:"type"`
	When    []string `json:"when"`
}
