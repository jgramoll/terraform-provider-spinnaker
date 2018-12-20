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

// MessageText for Pipeline Notification
type MessageText struct {
	Text string `json:"text"`
}

// PipelineCompleteKey for pipeline complete
const PipelineCompleteKey = "pipeline.complete"

// PipelineFailedKey for pipeline failed
const PipelineFailedKey = "pipeline.failed"

// PipelineStartingKey for pipeline starting
const PipelineStartingKey = "pipeline.starting"

// Message for Pipeline Notification
type Message struct {
	Complete MessageText `json:"pipeline.complete" mapstructure:"pipeline.complete"`
	Failed   MessageText `json:"pipeline.failed" mapstructure:"pipeline.failed"`
	Starting MessageText `json:"pipeline.starting" mapstructure:"pipeline.starting"`
}
